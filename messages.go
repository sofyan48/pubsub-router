package pubsubrouter

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
)

type Message struct {
	Attribute   map[string]string
	Payload     *pubsub.Message
	Data        []byte
	PublishTime time.Time
	CtlContext  context.Context
}
