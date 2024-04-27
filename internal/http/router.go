package http

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/docs"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/webserver/logger"
	"github.com/khivuksergey/webserver/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Router struct {
	*echo.Echo
}

func NewRouter(cfg *config.Configuration, services *service.Manager, logger logger.Logger) http.Handler {
	handlers := newHandlers(services, logger)

	e := router.NewEchoRouter().
		WithConfig(cfg.Router).
		UseMiddleware(handlers.error.HandleError).
		UseHealthCheck().
		UseSwagger(docs.SwaggerInfo, cfg.Swagger)

	// Login
	e.POST("/login", handlers.authorization.Login)

	// Users
	users := e.Group("/users")
	users.POST("", handlers.user.CreateUser)

	userId := users.Group("/:userId", handlers.authentication.AuthenticateJWT)
	userId.DELETE("", handlers.user.DeleteUser)
	userId.PUT("/username", handlers.user.UpdateUserName)
	userId.PUT("/password", handlers.user.UpdateUserPassword)

	return e
}
