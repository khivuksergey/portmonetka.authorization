package handler

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
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

// Health checks the health of the service.
//
// @Summary Check service health
// @Description Checks if the service is working properly
// @ID health
// @Produce json
// @Success 200 {object} common.Response "OK"
// @Router /health [get]
func (h HealthHandler) Health(c echo.Context) error {
	h.logger.Info(logger.LogMessage{
		Action:  "HealthHandler",
		Message: "Service is working",
	})
	return c.JSON(http.StatusOK, common.Response{Message: "OK"})
}
