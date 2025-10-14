package pubsubrouter

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub/v2"
)

type Message struct {
	ID          string
	Attribute   map[string]string
	Payload     *pubsub.Message
	Data        []byte
	PublishTime time.Time
	CtlContext  context.Context
}
