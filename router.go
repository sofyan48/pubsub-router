package pubsubrouter

import (
	"sync"

	"github.com/sofyan48/pubsub-router/pkg/client"
)

type Router struct {
	sync.Mutex
	handlers map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: map[string]Handler{},
	}
}

func (r *Router) Handle(routes string, h Handler) *Router {
	r.Lock()
	defer r.Unlock()

	r.handlers[routes] = h

	return r
}

func (r *Router) HandleMessage(m *Message) error {
	path := m.Payload.Attributes[client.MessageAttributeNameRoute]
	h, okRoute := r.handlers[path]
	if okRoute {
		err := h.HandleMessage(m)
		if err != nil {
			m.Payload.Nack()
			return err
		}
		m.Payload.Ack()
	}
	// if you need reporting please contrib this error handling
	// return errors.New("Route Not Any Match")
	return nil
}
