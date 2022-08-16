package pubsubrouter

import (
	"errors"
	"sync"
)

const (
	MessageAttributeNameRoute = "path"
)

type Router struct {
	sync.Mutex
	// routes   string
	// handle   handler.Handler
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
	path := m.Payload.Attributes[MessageAttributeNameRoute]
	h, okRoute := r.handlers[path]
	if okRoute {
		m.Payload.Ack()
		return h.HandleMessage(m)
	}
	m.Payload.Ack()
	return errors.New("route not any match")
}
