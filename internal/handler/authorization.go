package handler

import (
	"github.com/go-playground/validator/v10"
	serviceerror "github.com/khivuksergey/portmonetka.authorization/error"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"github.com/khivuksergey/portmonetka.common"
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
//	@Tags Authorization
//	@Summary		Login with user credentials
//	@Description	Login by username and password
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserLoginDTO	true 	"User login information"
//	@Success		200		{object}	model.Response
//	@Failure		400		{object}	model.Response "Bad Request: Invalid user data"
//	@Failure		401		{object}	model.Response "Unauthorized: Invalid credentials"
//	@Router			/login [post]
func (a AuthorizationHandler) Login(c echo.Context) error {
	requestUuid := c.Get(common.RequestUuidKey).(string)

	userLoginDTO := new(model.UserLoginDTO)
	err := a.bindUserLoginDtoValidate(c, userLoginDTO)
	if err != nil {
		a.logger.Error(logger.LogMessage{
			Action:      "Login",
			Message:     "Input data validation failed",
			Data:        err,
			RequestUuid: requestUuid,
		})
		return common.NewValidationError(serviceerror.InvalidInputData, err)
	}

	tokenResponse, err := a.authorizationService.Login(userLoginDTO)
	if err != nil {
		a.logger.Error(logger.LogMessage{
			Action:      "Login",
			Message:     "Login failed",
			Data:        err,
			RequestUuid: requestUuid,
		})
		return common.NewAuthorizationError(serviceerror.LoginFailed, err)
	}

	a.logger.Info(logger.LogMessage{
		Action:      "Login",
		Message:     "User logged in",
		UserId:      &userLoginDTO.Id,
		RequestUuid: requestUuid,
	})

	return c.JSON(http.StatusOK, model.Response{
		Message:     "logged in successfully",
		Data:        tokenResponse,
		RequestUuid: requestUuid,
	})
}

func (a AuthorizationHandler) bindUserLoginDtoValidate(c echo.Context, userLoginDTO *model.UserLoginDTO) error {
	if err := c.Bind(userLoginDTO); err != nil {
		return err
	}
	if err := a.validate.Struct(userLoginDTO); err != nil {
		return err
	}
	return nil
}
