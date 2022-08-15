package router

import (
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/sofyan48/pubsub-router/handler"
)

const (
	MessageAttributeNameRoute = "path"
)

type Router struct {
	sync.Mutex
	// routes   string
	// handle   handler.Handler
	handlers map[string]handler.Handler
	Resolver func(m *pubsub.Message) (string, string, bool)
}

func NewRouter() *Router {
	return &Router{
		handlers: map[string]handler.Handler{},
	}
}

func (r *Router) Handle(routes string, h handler.Handler) *Router {
	r.Lock()
	defer r.Unlock()

	r.handlers[routes] = h

	return r
}

func (r *Router) HandleMessage(m *pubsub.Message) error {
	path := m.Attributes[MessageAttributeNameRoute]

	h, okRoute := r.handlers[path]

	if okRoute {
		return h.HandleMessage(m)
	}

	m.Ack()
	return nil
}
