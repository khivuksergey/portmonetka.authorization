package http

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/handler"
	"github.com/khivuksergey/portmonetka.authorization/internal/handler/middleware"
	"github.com/khivuksergey/webserver/logger"
)

type Handlers struct {
	error          *middleware.ErrorHandlingMiddleware
	authentication *middleware.AuthenticationMiddleware
	authorization  *handler.AuthorizationHandler
	user           *handler.UserHandler
}

func newHandlers(services *service.Manager, logger logger.Logger) Handlers {
	return Handlers{
		error:          middleware.NewErrorHandlingMiddleware(),
		authentication: middleware.NewAuthenticationMiddleware(logger),
		authorization:  handler.NewAuthorizationHandler(services, logger),
		user:           handler.NewUserHandler(services, logger),
	}
}
