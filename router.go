package main

import (
	"github.com/khivuksergey/portmonetka.authorization/middleware"
	"github.com/khivuksergey/portmonetka.authorization/webservice"
	"github.com/khivuksergey/webserver"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func NewRouter(cfg *webserver.HttpHandlerConfig, webservice webservice.WebService) *echo.Echo {
	e := echo.New()

	if cfg != nil {
		if cfg.UseLogger {
			e.Use(echomiddleware.Logger())
		}
		if cfg.UseRecovery {
			e.Use(echomiddleware.Recover())
		}
	}

	e.GET("/health", webservice.Health)

	e.POST("/login", webservice.Login)

	e.POST("/users", webservice.CreateUser)
	e.DELETE("/users/:userId", webservice.DeleteUser, middleware.JWTAuthorization, middleware.Authentication)
	e.PUT("/users/:userId/username", webservice.UpdateUserName, middleware.JWTAuthorization, middleware.Authentication)
	e.PUT("/users/:userId/password", webservice.UpdateUserPassword, middleware.JWTAuthorization, middleware.Authentication)

	return e
}
