package discovery

import (
	"context"
)

type discoveryServerImpl struct {
	UnimplementedDiscoveryServer

	clientRegistry *ServerRegistry
}

func NewServer(cr *ServerRegistry) DiscoveryServer {
	return &discoveryServerImpl{
		clientRegistry: cr,
	}
}

func (d *discoveryServerImpl) ListGRPCServer(_ context.Context, _ *ListGRPCServerRequest) (*ListGRPCServerResponse, error) {
	return &ListGRPCServerResponse{
		Servers: d.clientRegistry.List(),
	}, nil
}

func (d *discoveryServerImpl) WatchGRPCServer(_ *WatchGRPCServerRequest, stream Discovery_WatchGRPCServerServer) error {
	obsChan := make(chan *WatchGRPCServerResponse)
	d.clientRegistry.AddObserver(obsChan)
	defer d.clientRegistry.RemoveObserver(obsChan)

	for {
		select {
		case event := <-obsChan:
			err := stream.Send(event)
			if err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
