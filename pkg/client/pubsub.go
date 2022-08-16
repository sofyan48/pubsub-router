package client

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/sofyan48/pubsub-router/pkg/session"
)

type clientPubsub struct {
	sessPubsub *pubsub.Client
	ctx        context.Context
}

const (
	MessageAttributeNameRoute = `path`
)

func NewClient(credential session.Contract) (*clientPubsub, error) {
	sess, err := pubsub.NewClient(credential.Context(), credential.GetConfig().ProjectID, credential.Option()...)
	if err != nil {
		return nil, err
	}
	return &clientPubsub{
		ctx:        credential.Context(),
		sessPubsub: sess,
	}, nil
}

func (c *clientPubsub) Client() *pubsub.Client {
	return c.sessPubsub
}
