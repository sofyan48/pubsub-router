package controller

import (
	pubsubrouter "github.com/sofyan48/pubsub-router"
)

type UseCase interface {
	Serve(m *pubsubrouter.Message) error
}
