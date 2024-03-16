package handler

import (
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthHandler struct {
	logger logger.Logger
}

func NewHealthHandler(logger logger.Logger) *HealthHandler {
	return &HealthHandler{
		logger: logger,
	}
}

func (h HealthHandler) Health(c echo.Context) error {
	h.logger.Info(logger.LogMessage{
		Action:  "HealthHandler",
		Message: "Service is working",
	})
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
