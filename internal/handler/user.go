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
// @Summary Create a new user
// @Description Creates a new user with the provided information
// @ID create-user
// @Accept json
// @Produce json
// @Param user body model.UserCreateDTO true "User object to be created"
// @Success 201 {object} common.Response "User created"
// @Failure 400 {object} common.Response "Bad request"
// @Failure 422 {object} common.Response "Unprocessable entity"
// @Router /users [post]
func (u UserHandler) CreateUser(c echo.Context) error {
	userCreateDTO, err := u.bindUserCreateDtoValidate(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	userId, err := u.userService.CreateUser(userCreateDTO)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.Response{
			Message: err.Error(),
		})
	}

	u.logger.Info(logger.LogMessage{
		Action:  "CreateUser",
		Message: "User created",
		UserId:  userId,
	})

	return c.JSON(http.StatusCreated, common.Response{
		Message: "User created",
		Data:    map[string]any{"userId": *userId},
	})
}

// DeleteUser deletes a user by ID.
//
// @Summary Delete user
// @Description Deletes user by the provided user ID
// @ID delete-user
// @Accept json
// @Produce json
// @Param userId path uint64 true "User ID"
// @Success 204 {string} string "No content"
// @Failure 400 {object} common.Response "Bad request"
// @Failure 422 {object} common.Response "Unprocessable entity"
// @Router /users/{userId} [delete]
func (u UserHandler) DeleteUser(c echo.Context) error {
	userId, ok := c.Get("userId").(uint64)
	if !ok {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: "invalid user id",
		})
	}

	if err := u.userService.DeleteUser(userId); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.Response{
			Message: err.Error(),
		})
	}

	u.logger.Info(logger.LogMessage{
		Action:  "DeleteUser",
		Message: "User deleted",
		UserId:  &userId,
	})

	return c.NoContent(http.StatusNoContent)
}

// UpdateUserName updates the name of a user.
//
// @Summary Update username
// @Description Updates the name of a user
// @ID update-user-name
// @Accept json
// @Produce json
// @Param userId path uint64 true "Authorized user ID"
// @Param user body model.UserUpdateNameDTO true "User update name request"
// @Success 200 {object} common.Response "Username updated"
// @Failure 400 {object} common.Response "Bad request"
// @Failure 401 {object} common.Response "Unauthorized"
// @Failure 422 {object} common.Response "Unprocessable entity"
// @Router /users/{userId}/username [put]
func (u UserHandler) UpdateUserName(c echo.Context) error {
	authorizedUserId, ok := c.Get("userId").(uint64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.Response{
			Message: "unauthorized user",
		})
	}

	userUpdateNameDTO, err := u.bindUserUpdateNameDtoValidate(c, authorizedUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	err = u.userService.UpdateUserName(userUpdateNameDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.Response{
			Message: err.Error(),
		})
	}

	u.logger.Info(logger.LogMessage{
		Action:  "UpdateUserName",
		Message: "Username updated",
		UserId:  &authorizedUserId,
	})

	return c.JSON(http.StatusOK, common.Response{
		Message: "Username updated",
	})
}

// UpdateUserPassword updates the password of a user.
//
// @Summary Update user password
// @Description Updates the password of a user
// @ID update-user-password
// @Accept json
// @Produce json
// @Param userId path uint64 true "Authorized user ID"
// @Param user body model.UserUpdatePasswordDTO true "User update password request"
// @Success 200 {object} common.Response "User password updated"
// @Failure 400 {object} common.Response "Bad request"
// @Failure 401 {object} common.Response "Unauthorized"
// @Failure 422 {object} common.Response "Unprocessable entity"
// @Router /users/{userId}/password [put]
func (u UserHandler) UpdateUserPassword(c echo.Context) error {
	authorizedUserId, ok := c.Get("userId").(uint64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.Response{
			Message: "unauthorized user",
		})
	}

	userUpdatePasswordDTO, err := u.bindUserUpdatePasswordDtoValidate(c, authorizedUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	err = u.userService.UpdateUserPassword(userUpdatePasswordDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.Response{
			Message: err.Error(),
		})
	}

	u.logger.Info(logger.LogMessage{
		Action:  "UpdateUserPassword",
		Message: "User password updated",
		UserId:  &authorizedUserId,
	})

	return c.JSON(http.StatusOK, common.Response{
		Message: "User password updated",
	})
}

func (u UserHandler) bindUserCreateDtoValidate(c echo.Context) (*model.UserCreateDTO, error) {
	userCreateDTO := new(model.UserCreateDTO)
	if err := c.Bind(userCreateDTO); err != nil {
		return nil, common.InvalidUserData
	}
	if err := u.validate.Struct(userCreateDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userCreateDTO, nil
}

func (u UserHandler) bindUserUpdateNameDtoValidate(c echo.Context, userId uint64) (*model.UserUpdateNameDTO, error) {
	userUpdateNameDTO := new(model.UserUpdateNameDTO)
	if err := c.Bind(userUpdateNameDTO); err != nil {
		return nil, common.InvalidUserData
	}
	if err := u.validate.Struct(userUpdateNameDTO); err != nil {
		return nil, common.InvalidUserData
	}
	userUpdateNameDTO.Id = userId
	return userUpdateNameDTO, nil
}

func (u UserHandler) bindUserUpdatePasswordDtoValidate(c echo.Context, userId uint64) (*model.UserUpdatePasswordDTO, error) {
	userUpdatePasswordDTO := new(model.UserUpdatePasswordDTO)
	if err := c.Bind(userUpdatePasswordDTO); err != nil {
		return nil, common.InvalidUserData
	}
	if err := u.validate.Struct(userUpdatePasswordDTO); err != nil {
		return nil, common.InvalidUserData
	}
	userUpdatePasswordDTO.Id = userId
	return userUpdatePasswordDTO, nil
}
