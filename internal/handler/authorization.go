package handler

import (
	"github.com/go-playground/validator/v10"
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
	validate             *validator.Validate
}

func NewAuthorizationHandler(services *service.Manager, logger logger.Logger) *AuthorizationHandler {
	return &AuthorizationHandler{
		authorizationService: services.Authorization,
		logger:               logger,
		validate:             validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Login returns authorization token for existing user
//
//	@Summary		Login with user credentials
//	@Description	Login by username and password
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserLoginDTO	true 	"user login information"
//	@Success		200		{object}	common.Response
//	@Failure		400		{object}	common.Response "Bad Request: Invalid user data"
//	@Failure		401		{object}	common.Response "Unauthorized: Invalid credentials"
//	@Router			/login [post]
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
	if err := a.validate.Struct(userLoginDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userLoginDTO, nil
}
