package webservice

import (
	"errors"
	serviceerrors "github.com/khivuksergey/portmonetka.authorization/errors"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"github.com/khivuksergey/portmonetka.authorization/utility"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ws *webservice) bindUserCreateDtoValidate(c echo.Context) *model.UserCreateDTO {
	userCreateDTO := new(model.UserCreateDTO)

	if err := c.Bind(userCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
		return nil
	}

	if userCreateDTO.Name == "" || userCreateDTO.Password == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Name and password cannot be empty",
		})
		return nil
	}

	// ignores deleted, but there's a unique constraint in db
	if ws.userRepository.Exists(userCreateDTO.Name) {
		c.JSON(http.StatusConflict, model.Response{
			Message: "User with this name already exists",
		})
		return nil
	}

	return userCreateDTO
}

func (ws *webservice) getNewUserName(c echo.Context) *string {
	userUpdateNameDTO := new(model.UserUpdateNameDTO)

	if err := c.Bind(userUpdateNameDTO); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
		return nil
	}

	if userUpdateNameDTO.Name == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Name cannot be empty",
		})
		return nil
	}

	// ignores deleted, but there's a unique constraint in db
	if ws.userRepository.Exists(userUpdateNameDTO.Name) {
		c.JSON(http.StatusConflict, model.Response{
			Message: "User with this name already exists",
		})
		return nil
	}

	return &userUpdateNameDTO.Name
}

func (ws *webservice) getNewUserPasswordHash(c echo.Context) *string {
	userUpdatePasswordDTO := new(model.UserUpdatePasswordDTO)

	if err := c.Bind(userUpdatePasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
		return nil
	}

	if userUpdatePasswordDTO.Password == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Password cannot be empty",
		})
		return nil
	}

	hashedPassword, err := utility.HashPassword(userUpdatePasswordDTO.Password)
	if err != nil {
		return nil
	}

	return &hashedPassword
}

func (ws *webservice) bindUserLoginDtoValidatePassword(c echo.Context) *model.UserLoginDTO {
	userLoginDTO := new(model.UserLoginDTO)

	if err := c.Bind(userLoginDTO); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
		return nil
	}

	if userLoginDTO.Name == "" || userLoginDTO.Password == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: "name and password cannot be empty",
		})
		return nil
	}

	user, err := ws.userRepository.FindUserByName(userLoginDTO.Name)
	if errors.Is(err, serviceerrors.UserNotFound) {
		c.JSON(http.StatusNotFound, model.Response{
			Message: err.Error(),
		})
		return nil
	}
	userLoginDTO.Id = user.Id

	if !utility.VerifyPassword(userLoginDTO.Password, user.Password) {
		c.JSON(http.StatusBadRequest, model.Response{
			Message: serviceerrors.InvalidPassword.Error(),
		})
		return nil
	}

	return userLoginDTO
}
