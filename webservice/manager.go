package webservice

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/db"
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type webservice struct {
	config         *config.Configuration
	logger         logger.Logger
	userRepository db.UserRepository
}

type WebService interface {
	Health(c echo.Context) error
	Login(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUserName(c echo.Context) error
	UpdateUserPassword(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func NewWebService(config *config.Configuration, logger logger.Logger) WebService {
	return &webservice{
		config:         config,
		logger:         logger,
		userRepository: db.NewUserRepository(config.DB),
	}
}

func (ws *webservice) Health(c echo.Context) error {
	ws.logger.Info(logger.LogMessage{
		Action:  "Health",
		Message: "Service is working",
	})
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}
