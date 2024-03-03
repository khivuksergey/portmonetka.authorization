package main

import (
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type webservice struct {
	logger logger.Logger
}

type WebService interface {
	Health(c echo.Context) error
}

func NewWebService(logger logger.Logger) WebService {
	return &webservice{logger: logger}
}

func (ws *webservice) Health(c echo.Context) error {
	ws.logger.Info(logger.LogMessage{
		Action:  "Health",
		Message: "Service is working",
	})
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
