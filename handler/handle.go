package handler

import "cloud.google.com/go/pubsub"

type Handler interface {
	HandleMessage(*pubsub.Message) error
}

// HandlerFunc is an adaptor to allow the use of ordinary functions as message Handlers.
type HandlerFunc func(*pubsub.Message) error

// HandleMessage satisfies the Handler interface.
func (h HandlerFunc) HandleMessage(m *pubsub.Message) error {
	return h(m)
}
