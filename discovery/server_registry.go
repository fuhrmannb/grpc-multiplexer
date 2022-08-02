package discovery

import (
	"fmt"
	"sync"

	"google.golang.org/grpc"
)

type ServerRegistry struct {
	// Use it as map[string]*grpc.ClientConn
	streamManagers sync.Map
	// Use if as map[chan *WatchGRPCServerResponse]bool
	observers sync.Map
}

func NewServerRegistry() *ServerRegistry {
	return &ServerRegistry{}
}

func (cr *ServerRegistry) Register(server string, stream *grpc.ClientConn) error {
	if _, exists := cr.Get(server); exists {
		return fmt.Errorf("server %v already registered", server)
	}
	cr.streamManagers.Store(server, stream)

	cr.notifyObservers(server, WatchGRPCServerResponse_EVENT_CONNECTED)

	return nil
}

func (cr *ServerRegistry) Unregister(server string) {
	cr.streamManagers.Delete(server)

	cr.notifyObservers(server, WatchGRPCServerResponse_EVENT_DISCONNECTED)
}

func (cr *ServerRegistry) Get(server string) (*grpc.ClientConn, bool) {
	stream, ok := cr.streamManagers.Load(server)
	if !ok {
		return nil, ok
	}
	return stream.(*grpc.ClientConn), ok
}

func (cr *ServerRegistry) List() []string {
	var servers []string

	cr.streamManagers.Range(func(k interface{}, v interface{}) bool {
		servers = append(servers, k.(string))
		return true
	})

	return servers
}

func (cr *ServerRegistry) AddObserver(c chan *WatchGRPCServerResponse) {
	cr.observers.Store(c, true)
}

func (cr *ServerRegistry) RemoveObserver(o chan *WatchGRPCServerResponse) {
	cr.observers.Delete(o)
}

func (cr *ServerRegistry) notifyObservers(server string, event WatchGRPCServerResponse_Event) {
	cr.observers.Range(func(o, _ interface{}) bool {
		observer := o.(chan *WatchGRPCServerResponse)
		go func() {
			observer <- &WatchGRPCServerResponse{
				Server: server,
				Event:  event,
			}
		}()
		return true
	})
}
