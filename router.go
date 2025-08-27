package pubsubrouter

import (
	"runtime/debug"
	"sync"

	"github.com/google/martian/log"
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
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("panic recovered: %v | stack : %v", err, string(debug.Stack()))
		}
	}()
	h, okRoute := r.handlers[path]
	if okRoute {
		m.Payload.Ack()
		return h.HandleMessage(m)
	}
	// if you need reporting please contrib this error handling
	// return errors.New("Route Not Any Match")
	return nil
}
