package multiplexer

import (
	"bufio"
	"context"
	"net"
	"strings"

	"github.com/fuhrmannb/grpc-proxy/proxy"
	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/fuhrmannb/grpc-multiplexer/discovery"
)

type server struct {
	proxyLis       net.Listener
	multiplexerLis net.Listener
	serverRegistry *discovery.ServerRegistry
}

func NewServer(proxyLis net.Listener, multiplexerLis net.Listener) *server {
	return &server{
		proxyLis:       proxyLis,
		multiplexerLis: multiplexerLis,
		serverRegistry: discovery.NewServerRegistry(),
	}
}

func (s *server) Run() error {
	// Listen for multiplexer client connections
	go func() {
		for {
			incoming, err := s.multiplexerLis.Accept()
			if err != nil {
				log.Error().Err(err).Msg("couldn't accept multiplexer TCP connection")
				continue
			}
			logger := log.With().Str("client_addr", incoming.RemoteAddr().String()).Logger()
			logger.Debug().Msg("accept new multiplexer TCP connection")
			go func() {
				err := s.handleClient(incoming, &logger)
				if err != nil {
					logger.Err(err).Msg("error handling client")
				}
			}()
		}
	}()

	// Start Proxy server with request filtering by clientID
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		clientID, outCtx, err := extractClientIDNameFromCtx(ctx)
		if err != nil {
			return nil, nil, err
		}

		clientConn, ok := s.serverRegistry.Get(clientID)
		if !ok {
			return nil, nil, status.Errorf(codes.NotFound, "Client '%v' not registered", clientID)
		}
		return outCtx, clientConn, nil
	}
	// FIXME: use encoding.RegisterCodec + add Name() method to proxy.Codec
	//nolint:staticcheck
	proxyServer := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))
	discovery.RegisterDiscoveryServiceServer(proxyServer, discovery.NewServer(s.serverRegistry))
	logger := log.With().Str("proxy_lis_addr", s.proxyLis.Addr().String()).Str("multiplexer_lis_addr", s.multiplexerLis.Addr().String()).Logger()
	logger.Info().Msg("ready to listen to new connection")
	err := proxyServer.Serve(s.proxyLis)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) handleClient(clientConn net.Conn, logger *zerolog.Logger) error {
	// Accept initial Yamux connection from client
	initSession, err := yamux.Server(clientConn, nil)
	if err != nil {
		return errors.Wrap(err, "can't create yamux server")
	}
	defer initSession.Close()
	logger.Debug().Msg("yamux server created")
	initStream, err := initSession.Accept()
	if err != nil {
		return errors.Wrap(err, "can't accept yamux server")
	}
	logger.Debug().Msg("yamux session accepted")

	// Read ClientID from client
	serverReader := bufio.NewReader(initStream)
	clientID, err := serverReader.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "can't read client ID")
	}
	clientID = strings.TrimSuffix(clientID, "\n")
	*logger = logger.With().Str("client_id", clientID).Logger()
	logger.Debug().Msg("client ID received")

	// Create a session with client to retrieve gRPC connection
	gRPCSession, err := yamux.Client(initStream, yamux.DefaultConfig())
	if err != nil {
		return errors.Wrap(err, "can't open yamux client")
	}
	defer gRPCSession.Close()
	logger.Debug().Msg("yamux client opened")
	gRPCStream, err := gRPCSession.Open()
	if err != nil {
		return errors.Wrap(err, "can't open yamux session for gRPC")
	}
	defer gRPCStream.Close()
	logger.Debug().Msg("yamux session for gRPC opened")

	// Get gRPCClient from Yamux connection
	// Note: no target, using Yamux stream as gRPC connection
	gRPCClient, err := grpc.Dial("", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return gRPCStream, nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "can't gRPC dial")
	}
	logger.Debug().Msg("gRPC dial done")

	// Register client
	err = s.serverRegistry.Register(clientID, gRPCClient)
	if err != nil {
		return errors.Wrap(err, "can't gRPC dial")
	}
	defer s.serverRegistry.Unregister(clientID)
	logger.Info().Msg("client connected")

	// Wait for client connection or end of the program
	<-initSession.CloseChan()
	<-gRPCSession.CloseChan()

	logger.Info().Msg("client disconnected")
	return nil
}

const ClientIDMDKeyName = "client-id"

func clientIDFromMetadata(md metadata.MD) (string, error) {
	names, ok := md[ClientIDMDKeyName]
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "no metadata with key '%v' supplied", ClientIDMDKeyName)
	}
	if len(names) != 1 {
		return "", status.Errorf(codes.InvalidArgument, "request needs only one element for metadata '%v'", ClientIDMDKeyName)
	}

	return names[0], nil
}

func extractClientIDNameFromCtx(ctx context.Context) (string, context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", nil, status.Errorf(codes.Internal, "can't retrieve metadata from context")
	}

	name, err := clientIDFromMetadata(md)
	if err != nil {
		return "", nil, err
	}

	outMD := md.Copy()
	delete(outMD, ClientIDMDKeyName)

	return name, metadata.NewOutgoingContext(ctx, outMD), nil
}
