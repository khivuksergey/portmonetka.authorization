package handler

import (
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
}

func NewUserHandler(services *service.Manager, logger logger.Logger) *UserHandler {
	return &UserHandler{
		userService: services.User,
		logger:      logger,
	}
}

func (u UserHandler) CreateUser(c echo.Context) error {
	userCreateDTO, err := u.bindUserCreateDto(c)
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

// DeleteUser is authorized by custom middleware
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

// UpdateUserName is authorized by custom middleware
func (u UserHandler) UpdateUserName(c echo.Context) error {
	userUpdateNameDTO, err := u.bindUserUpdateNameDto(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	authorizedUserId, ok := c.Get("userId").(uint64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.Response{
			Message: "unauthorized user",
		})
	}

	userUpdateNameDTO.Id = authorizedUserId
	err = u.userService.UpdateUserName(userUpdateNameDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, common.Response{
			Message: err.Error(),
		})
	}

	u.logger.Info(logger.LogMessage{
		Action:  "UpdateUserName",
		Message: "User name updated",
		UserId:  &authorizedUserId,
	})

	return c.JSON(http.StatusOK, common.Response{
		Message: "User name updated",
	})
}

// UpdateUserPassword is authorized by custom middleware
func (u UserHandler) UpdateUserPassword(c echo.Context) error {
	userUpdatePasswordDTO, err := u.bindUserUpdatePasswordDto(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Message: err.Error(),
		})
	}

	authorizedUserId, ok := c.Get("userId").(uint64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.Response{
			Message: "unauthorized user",
		})
	}

	userUpdatePasswordDTO.Id = authorizedUserId
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

func (u UserHandler) bindUserCreateDto(c echo.Context) (*model.UserCreateDTO, error) {
	userCreateDTO := new(model.UserCreateDTO)
	if err := c.Bind(userCreateDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userCreateDTO, nil
}

func (u UserHandler) bindUserUpdateNameDto(c echo.Context) (*model.UserUpdateNameDTO, error) {
	userUpdateNameDTO := new(model.UserUpdateNameDTO)
	if err := c.Bind(userUpdateNameDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userUpdateNameDTO, nil
}

func (u UserHandler) bindUserUpdatePasswordDto(c echo.Context) (*model.UserUpdatePasswordDTO, error) {
	userUpdatePasswordDTO := new(model.UserUpdatePasswordDTO)
	if err := c.Bind(userUpdatePasswordDTO); err != nil {
		return nil, common.InvalidUserData
	}
	return userUpdatePasswordDTO, nil
}
