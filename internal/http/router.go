package http

import (
	//_ "github.com/khivuksergey/portmonetka.authorization/docs"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	*echo.Echo
}

func NewRouter(cfg *webserver.RouterConfig, services *service.Manager, logger logger.Logger) *echo.Echo {
	e := echo.New()

	if cfg != nil {
		if cfg.UseLogger {
			e.Use(middleware.Logger())
		}
		if cfg.UseRecovery {
			e.Use(middleware.Recover())
		}
	}

	handlers := newHandlers(services, logger)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/health", handlers.health.Health)

	e.POST("/login", handlers.authorization.Login)

	usersGroup := e.Group("/users")
	usersGroup.POST("", handlers.user.CreateUser)

	userRoutes := usersGroup.Group("/:userId")
	userRoutes.Use(handlers.middleware.JWT, handlers.middleware.Authentication)

	userRoutes.DELETE("", handlers.user.DeleteUser)
	userRoutes.PUT("/username", handlers.user.UpdateUserName)
	userRoutes.PUT("/password", handlers.user.UpdateUserPassword)

	return e
}
