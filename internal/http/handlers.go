package http

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/handler"
	"github.com/khivuksergey/webserver/logger"
)

type Handlers struct {
	middleware    *handler.Middleware
	health        *handler.HealthHandler
	authorization *handler.AuthorizationHandler
	user          *handler.UserHandler
}

func newHandlers(services *service.Manager, logger logger.Logger) Handlers {
	return Handlers{
		middleware:    handler.NewMiddleware(logger),
		health:        handler.NewHealthHandler(logger),
		authorization: handler.NewAuthorizationHandler(services, logger),
		user:          handler.NewUserHandler(services, logger),
	}
}
