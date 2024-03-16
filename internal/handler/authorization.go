package handler

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthorizationHandler struct {
	authorizationService service.AuthorizationService
	logger               logger.Logger
}

func NewAuthorizationHandler(services *service.Manager, logger logger.Logger) *AuthorizationHandler {
	return &AuthorizationHandler{
		authorizationService: services.Authorization,
		logger:               logger,
	}
}

func (a AuthorizationHandler) Login(c echo.Context) error {
	userLoginDTO, err := a.bindUserLoginDto(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	tokenResponse, err := a.authorizationService.Login(userLoginDTO)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, common.Response{
			Message: err.Error(),
		})
	}

	a.logger.Info(logger.LogMessage{
		Action:  "Login",
		Message: "User logged in",
		UserId:  &userLoginDTO.Id,
	})

	return c.JSON(http.StatusOK, common.Response{
		Message: "logged in successfully",
		Data:    tokenResponse,
	})
}

func (a AuthorizationHandler) bindUserLoginDto(c echo.Context) (*model.UserLoginDTO, error) {
	userLoginDTO := new(model.UserLoginDTO)
	if err := c.Bind(userLoginDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userLoginDTO, nil
}
