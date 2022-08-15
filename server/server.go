package server

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/sofyan48/pubsub-router/pkg/client"
	"github.com/sofyan48/pubsub-router/pkg/session"
	"github.com/sofyan48/pubsub-router/router"
)

type Server struct {
	clients   *pubsub.Client
	ctx       context.Context
	subClient *pubsub.Subscription
	router    *router.Router
}

func NewServer(ctx context.Context, sess session.Contract) *Server {
	cl, err := client.NewClient(sess)
	if err != nil {
		defer cl.Client().Close()
	}
	return &Server{
		clients: cl.Client(),
		ctx:     ctx,
	}
}

func (s *Server) Subscribe(topic string, r *router.Router) *Server {
	s.subClient = s.clients.Subscription(topic)
	s.router = r
	return s
}

func (s *Server) Start() {
	s.subClient.Receive(s.ctx, func(ctx context.Context, m *pubsub.Message) {
		s.router.HandleMessage(m)
	})
}
