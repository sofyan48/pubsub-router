package pubsubrouter

import (
	"errors"
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
		m.Payload.Ack()
		return h.HandleMessage(m)
	}
	return errors.New("route not any match")
}
