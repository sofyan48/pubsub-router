package pubsubrouter

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/sofyan48/pubsub-router/pkg/client"
	"github.com/sofyan48/pubsub-router/pkg/session"
)

type Server struct {
	clients   *pubsub.Client
	ctx       context.Context
	subClient *pubsub.Subscription
	router    *Router
}

func NewSession(ctx context.Context, sess session.Contract) *Server {
	cl, err := client.NewClient(sess)
	if err != nil {
		defer cl.Client().Close()
	}
	return &Server{
		clients: cl.Client(),
		ctx:     ctx,
	}
}

func (s *Server) Subscribe(topic string, r *Router) *Server {
	s.subClient = s.clients.Subscription(topic)
	s.router = r
	return s
}

func (s *Server) Publish(topic, path, msg string) (string, error) {
	if path == "" {
		return "", errors.New("path is required")
	}
	cl := s.clients.Topic(topic)
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

func (s *Server) Start() {
	var received int32
	s.subClient.Receive(s.ctx, func(ctx context.Context, msg *pubsub.Message) {
		atomic.AddInt32(&received, 1)
		m := Message{}
		m.Data = msg.Data
		m.Attribute = msg.Attributes
		m.Payload = msg
		m.PublishTime = msg.PublishTime
		m.CtlContext = s.ctx
		m.ID = msg.ID
		err := s.router.HandleMessage(&m)
		if err != nil {
			msg.Ack()
			fmt.Println("error", err.Error())
		}
	})
}
