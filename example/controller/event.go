package controller

import (
	"fmt"

	pubsubrouter "github.com/sofyan48/pubsub-router"
)

type event struct {
}

func NewEvent() UseCase {
	return &event{}
}

// Serve implements UseCase
func (*event) Serve(m *pubsubrouter.Message) error {
	fmt.Println(string(m.Data))
	return nil
}
