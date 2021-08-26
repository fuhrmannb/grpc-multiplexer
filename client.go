package multiplexer

import (
	"net"
	"time"

	"github.com/fuhrmannb/grpc-proxy/proxy"
	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	gRPCServerAddr  string
	multiplexerAddr string
	clientID        string
}

func NewClient(clientID string, gRPCServerAddr string, multiplexerAddr string) *client {
	return &client{
		clientID:        clientID,
		gRPCServerAddr:  gRPCServerAddr,
		multiplexerAddr: multiplexerAddr,
	}
}

func (c *client) Run() error {
	logger := log.With().Str("multiplexer_addr", c.multiplexerAddr).Logger()

	multiplexerConn, err := net.DialTimeout("tcp", c.multiplexerAddr, time.Second*10)
	if err != nil {
		return errors.Wrap(err, "can't connect to multiplexer")
	}
	logger.Debug().Msg("connection created to multiplexer")
	defer multiplexerConn.Close()
	yamuxClient, err := yamux.Client(multiplexerConn, nil)
	if err != nil {
		return errors.Wrap(err, "can't create yamux client")
	}
	logger.Debug().Msg("yamux client created")
	yamuxConn, err := yamuxClient.Open()
	if err != nil {
		return errors.Wrap(err, "can't open yamux connection")
	}
	defer yamuxConn.Close()
	logger.Debug().Msg("yamux client opened")

	logger = logger.With().Str("client_id", c.clientID).Logger()
	_, err = yamuxConn.Write(append([]byte(c.clientID), '\n'))
	if err != nil {
		return errors.Wrap(err, "can't send client ID")
	}
	logger.Debug().Msg("client ID sent to multiplexer")

	lis, err := yamux.Server(yamuxConn, yamux.DefaultConfig())
	if err != nil {
		return errors.Wrap(err, "can't create yamux server")
	}
	defer lis.Close()
	logger.Debug().Msg("client: yamux server created")

	logger.Info().Msg("connected to multiplexer")

	logger = logger.With().Str("grpc_server_addr", c.gRPCServerAddr).Logger()
	clientConn, err := grpc.Dial(c.gRPCServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "can't connect to gRPC Server")
	}
	logger.Info().Msg("connected to gRPC server")
	proxy := proxy.NewProxy(clientConn)
	logger.Info().Msg("serving proxy to forward requests...")
	err = proxy.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
