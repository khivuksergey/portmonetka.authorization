package webservice

import (
	"errors"
	"fmt"
	serviceerrors "github.com/khivuksergey/portmonetka.authorization/errors"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ws *webservice) Login(c echo.Context) error {
	userLoginDTO := ws.bindUserLoginDtoValidatePassword(c)
	if userLoginDTO == nil {
		return serviceerrors.UserDataValidationError
	}

	tokenResponse, err := ws.getToken(userLoginDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: fmt.Sprintf("get token error: %v", err),
		})
	}

	ws.userRepository.UpdateLastLoginTime(userLoginDTO.Id, tokenResponse.IssuedAt)

	return c.JSON(http.StatusOK, tokenResponse)
}

func (ws *webservice) CreateUser(c echo.Context) error {
	userCreateDTO := ws.bindUserCreateDtoValidate(c)
	if userCreateDTO == nil {
		return serviceerrors.UserDataValidationError
	}

	id, err := ws.userRepository.CreateUser(userCreateDTO.Name, userCreateDTO.Password)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, model.Response{
		Message: "Created user",
		Data:    map[string]any{"userId": id},
	})
}

// DeleteUser is authorized by custom middleware
func (ws *webservice) DeleteUser(c echo.Context) error {
	userId := c.Get("userId").(uint64)

	err := ws.userRepository.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: err.Error(),
		})
	}

	ws.logger.Info(logger.LogMessage{
		Action:  "DeleteUser",
		Message: "User deleted",
		UserId:  &userId,
	})

	return c.NoContent(http.StatusNoContent)
}

// UpdateUserName is authorized by custom middleware
func (ws *webservice) UpdateUserName(c echo.Context) error {
	newUserName := ws.getNewUserName(c)
	if newUserName == nil {
		return serviceerrors.UserDataValidationError
	}

	userIdToUpdate := c.Get("userId").(uint64)
	userToUpdate, err := ws.userRepository.FindUserById(userIdToUpdate)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Message: err.Error(),
		})
	}

	userToUpdate.Name = *newUserName

	err = ws.userRepository.UpdateUser(*userToUpdate)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: err.Error(),
		})
	}

	ws.logger.Info(logger.LogMessage{
		Action:  "UpdateUserName",
		Message: "User name updated",
		UserId:  &userIdToUpdate,
	})

	return c.JSON(http.StatusOK, model.Response{
		Message: "updated user name",
	})
}

// UpdateUserPassword is authorized by custom middleware
func (ws *webservice) UpdateUserPassword(c echo.Context) error {
	newUserPasswordHash := ws.getNewUserPasswordHash(c)
	if newUserPasswordHash == nil {
		return errors.New("user data validation error")
	}

	userIdToUpdate := c.Get("userId").(uint64)
	userToUpdate, err := ws.userRepository.FindUserById(userIdToUpdate)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Message: err.Error(),
		})
	}

	userToUpdate.Password = *newUserPasswordHash

	err = ws.userRepository.UpdateUser(*userToUpdate)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: err.Error(),
		})
	}

	ws.logger.Info(logger.LogMessage{
		Action:  "UpdateUserPassword",
		Message: "User password updated",
		UserId:  &userIdToUpdate,
	})

	return c.JSON(http.StatusOK, model.Response{
		Message: "updated user password",
	})
}
