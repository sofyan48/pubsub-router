package controller

import (
	"fmt"

	pubsubrouter "github.com/sofyan48/pubsub-router"
)

type test struct {
}

func Newtest() UseCase {
	return &event{}
}

// Serve implements UseCase
func (t *test) Serve(m *pubsubrouter.Message) error {
	fmt.Println(string(m.Data))
	return nil
}
