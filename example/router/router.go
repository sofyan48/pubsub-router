package router

import (
	pubsubrouter "github.com/sofyan48/pubsub-router"
	"github.com/sofyan48/pubsub-router/example/controller"
)

type Router struct {
	rtr *pubsubrouter.Router
}

func NewRouter() *Router {
	return &Router{
		rtr: pubsubrouter.NewRouter(),
	}
}

func (r *Router) handle(svc controller.UseCase) pubsubrouter.HandlerFunc {
	return svc.Serve
}

func (r *Router) Route() *pubsubrouter.Router {

	r.rtr.Handle("/event", r.handle(controller.NewEvent()))
	r.rtr.Handle("/test", r.handle(controller.Newtest()))

	return r.rtr
}
