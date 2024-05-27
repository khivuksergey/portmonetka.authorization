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

type UserHandler struct {
	userService service.UserService
	logger      logger.Logger
	validate    *validator.Validate
}

func NewUserHandler(services *service.Manager, logger logger.Logger) *UserHandler {
	return &UserHandler{
		userService: services.User,
		logger:      logger,
		validate:    validator.New(validator.WithRequiredStructEnabled()),
	}
}

// CreateUser creates a new user.
//
// @Tags User
// @Summary Create a new user
// @Description Creates a new user with the provided information
// @ID create-user
// @Accept json
// @Produce json
// @Param user body model.UserCreateDTO true "User object to be created"
// @Success 201 {object} model.Response "User created"
// @Failure 400 {object} model.Response "Bad request"
// @Failure 422 {object} model.Response "Unprocessable entity"
// @Router /users [post]
func (u UserHandler) CreateUser(c echo.Context) error {
	requestUuid := c.Get(common.RequestUuidKey).(string)
	userCreateDTO := &model.UserCreateDTO{}

	err := bindDtoValidate[model.UserCreateDTO](c, u.validate, userCreateDTO)
	if err != nil {
		return common.NewValidationError(serviceerror.InvalidInputData, err)
	}

	userId, err := u.userService.CreateUser(*userCreateDTO)
	if err != nil {
		return common.NewUnprocessableEntityError(serviceerror.CannotCreateUser, err)
	}

	u.logger.Info(logger.LogMessage{
		Action:      "CreateUser",
		Message:     "User created",
		UserId:      &userId,
		RequestUuid: requestUuid,
	})

	return c.JSON(http.StatusCreated, model.Response{
		Message:     "User created",
		Data:        map[string]any{"userId": userId},
		RequestUuid: requestUuid,
	})
}

// UpdateUserName updates the name of a user.
//
// @Tags User
// @Summary Update username
// @Description Updates the name of a user
// @ID update-user-name
// @Accept json
// @Produce json
// @Param userId path uint64 true "Authorized user ID"
// @Param user body model.UserUpdateNameDTO true "User update name request"
// @Success 200 {object} model.Response "Username updated"
// @Failure 400 {object} model.Response "Bad request"
// @Failure 401 {object} model.Response "Unauthorized"
// @Failure 422 {object} model.Response "Unprocessable entity"
// @Router /users/{userId}/username [put]
func (u UserHandler) UpdateUserName(c echo.Context) error {
	requestUuid := c.Get(common.RequestUuidKey).(string)
	userId := c.Get("userId").(uint64)
	userUpdateNameDTO := &model.UserUpdateNameDTO{}

	err := bindDtoValidate[model.UserUpdateNameDTO](c, u.validate, userUpdateNameDTO)
	if err != nil {
		return common.NewValidationError(serviceerror.InvalidInputData, err)
	}

	userUpdateNameDTO.Id = userId

	err = u.userService.UpdateUserName(*userUpdateNameDTO)
	if err != nil {
		return common.NewUnprocessableEntityError(serviceerror.CannotUpdateUsername, err)
	}

	u.logger.Info(logger.LogMessage{
		Action:      "UpdateUserName",
		Message:     "Username updated",
		UserId:      &userId,
		RequestUuid: requestUuid,
	})

	return c.JSON(http.StatusOK, model.Response{
		Message:     "Username updated",
		Data:        map[string]any{"name": userUpdateNameDTO.Name},
		RequestUuid: requestUuid,
	})
}

// UpdateUserPassword updates the password of a user.
//
// @Tags User
// @Summary Update user password
// @Description Updates the password of a user
// @ID update-user-password
// @Accept json
// @Produce json
// @Param userId path uint64 true "Authorized user ID"
// @Param user body model.UserUpdatePasswordDTO true "User update password request"
// @Success 200 {object} model.Response "User password updated"
// @Failure 400 {object} model.Response "Bad request"
// @Failure 401 {object} model.Response "Unauthorized"
// @Failure 422 {object} model.Response "Unprocessable entity"
// @Router /users/{userId}/password [put]
func (u UserHandler) UpdateUserPassword(c echo.Context) error {
	requestUuid := c.Get(common.RequestUuidKey).(string)
	userId := c.Get("userId").(uint64)
	userUpdatePasswordDTO := &model.UserUpdatePasswordDTO{}

	err := bindDtoValidate[model.UserUpdatePasswordDTO](c, u.validate, userUpdatePasswordDTO)
	if err != nil {
		return common.NewValidationError(serviceerror.InvalidInputData, err)
	}

	userUpdatePasswordDTO.Id = userId

	err = u.userService.UpdateUserPassword(*userUpdatePasswordDTO)
	if err != nil {
		return common.NewUnprocessableEntityError(serviceerror.CannotUpdatePassword, err)
	}

	u.logger.Info(logger.LogMessage{
		Action:      "UpdateUserPassword",
		Message:     "User password updated",
		UserId:      &userId,
		RequestUuid: requestUuid,
	})

	return c.JSON(http.StatusOK, model.Response{
		Message:     "User password updated",
		RequestUuid: requestUuid,
	})
}

// DeleteUser deletes a user by ID.
//
// @Tags User
// @Summary Delete user
// @Description Deletes user by the provided user ID
// @ID delete-user
// @Accept json
// @Produce json
// @Param userId path uint64 true "User ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} model.Response "Bad request"
// @Failure 422 {object} model.Response "Unprocessable entity"
// @Router /users/{userId} [delete]
func (u UserHandler) DeleteUser(c echo.Context) error {
	requestUuid := c.Get(common.RequestUuidKey).(string)
	userId := c.Get("userId").(uint64)

	if err := u.userService.DeleteUser(userId); err != nil {
		return common.NewUnprocessableEntityError(serviceerror.CannotDeleteUser, err)
	}

	u.logger.Info(logger.LogMessage{
		Action:      "DeleteUser",
		Message:     "User deleted",
		UserId:      &userId,
		RequestUuid: requestUuid,
	})

	return c.NoContent(http.StatusNoContent)
}

func bindDtoValidate[T any](c echo.Context, validate *validator.Validate, dto *T) error {
	if err := c.Bind(dto); err != nil {
		return err
	}
	if err := validate.Struct(dto); err != nil {
		return err
	}
	return nil
}
