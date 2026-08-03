package pubsubrouter

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"sync"

	"github.com/sofyan48/pubsub-router/pkg/client"
)

var ErrRouteNotFound = errors.New("route not found")

type Router struct {
	sync.RWMutex
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

func (r *Router) HandleMessage(m *Message) (err error) {
	if m == nil || m.Payload == nil {
		return errors.New("message payload is required")
	}

	path := m.Payload.Attributes[client.MessageAttributeNameRoute]
	defer func() {
		if recovered := recover(); recovered != nil {
			log.Printf("panic recovered: %v | stack: %s", recovered, debug.Stack())
			err = fmt.Errorf("handler panic recovered: %v", recovered)
		}
	}()

	r.RLock()
	h, okRoute := r.handlers[path]
	r.RUnlock()
	if !okRoute {
		return fmt.Errorf("%w: %s", ErrRouteNotFound, path)
	}

	return h.HandleMessage(m)
}
