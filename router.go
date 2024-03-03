package main

import (
	"github.com/khivuksergey/webserver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(config *webserver.HttpHandlerConfig, webservice WebService) *echo.Echo {
	e := echo.New()

	if config != nil {
		if config.UseLogger {
			e.Use(middleware.Logger())
		}
		if config.UseRecovery {
			e.Use(middleware.Recover())
		}
	}

	e.GET("/health", webservice.Health)

	return e
}
