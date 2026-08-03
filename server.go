package pubsubrouter

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/pubsub/v2"
	"github.com/sofyan48/pubsub-router/pkg/client"
	"github.com/sofyan48/pubsub-router/pkg/session"
)

type Server struct {
	clients   *pubsub.Client
	ctx       context.Context
	subClient *pubsub.Subscriber
	router    *Router
}

func NewSession(ctx context.Context, sess session.Contract) *Server {
	cl, err := client.NewClient(sess)
	if err != nil {
		log.Fatalf("pubsubrouter client not connected: %v", err)
	}
	return &Server{
		clients: cl.Client(),
		ctx:     ctx,
	}
}

func NewSessionAutoConfig(ctx context.Context, projectID string) *Server {
	cl, err := client.NewClientAutoConfig(ctx, projectID)
	if err != nil {
		log.Fatalf("pubsubrouter client not connected: %v", err)
	}
	return &Server{
		clients: cl.Client(),
		ctx:     ctx,
	}
}

func (s *Server) Subscribe(topic string, r *Router) *Server {
	s.subClient = s.clients.Subscriber(topic)
	s.router = r
	return s
}

func (s *Server) Publish(topic, path, msg string) (string, error) {
	if path == "" {
		return "", errors.New("path is required")
	}
	cl := s.clients.Publisher(topic)
	cl.PublishSettings.NumGoroutines = 1
	return cl.Publish(
		s.ctx,
		&pubsub.Message{
			Data:        []byte(msg),
			Attributes:  map[string]string{client.MessageAttributeNameRoute: path},
			PublishTime: time.Now(),
		},
	).Get(s.ctx)
}

func (s *Server) Start() error {
	if s.subClient == nil {
		return errors.New("subscriber is not configured")
	}
	if s.router == nil {
		return errors.New("router is not configured")
	}

	return s.subClient.Receive(s.ctx, func(ctx context.Context, msg *pubsub.Message) {
		m := Message{
			Data:        msg.Data,
			Attribute:   msg.Attributes,
			Payload:     msg,
			PublishTime: msg.PublishTime,
			CtlContext:  ctx,
			ID:          msg.ID,
		}
		if err := s.router.HandleMessage(&m); err != nil {
			msg.Nack()
			return
		}
		msg.Ack()
	})
}
