package main

import (
	"errors"
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"github.com/khivuksergey/portmonetka.authorization/utility"
	"github.com/khivuksergey/webserver/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var secretKey = []byte("'b^i['Qs{qQOh\"_`9%jU1XVzAdx89gF(p)14L8DW{Zv.jwM[Ue3aj!w=iWsNk$I")

type webservice struct {
	config         *config.Configuration
	logger         logger.Logger
	userRepository UserRepository
}

type WebService interface {
	Health(c echo.Context) error
	Login(c echo.Context) error
	CreateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func NewWebService(config *config.Configuration, logger logger.Logger) WebService {
	return &webservice{
		config:         config,
		logger:         logger,
		userRepository: NewUserRepository(config.DB),
	}
}

func (ws *webservice) Health(c echo.Context) error {
	ws.logger.Info(logger.LogMessage{
		Action:  "Health",
		Message: "Service is working",
	})
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func (ws *webservice) Login(c echo.Context) error {
	var userDTO model.UserDTO

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
	}

	if userDTO.Name == "" || userDTO.Password == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: "name and password cannot be empty",
		})
	}

	user, err := ws.userRepository.FindUserByName(userDTO.Name)
	if errors.Is(err, UserNotFound) {
		return c.JSON(http.StatusNotFound, model.Response{
			Message: err.Error(),
		})
	}

	if !utility.VerifyPassword(userDTO.Password, user.Password) {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: InvalidPassword.Error(),
		})
	}

	user.RememberMe = userDTO.RememberMe

	tokenResponse, err := ws.getToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: fmt.Sprintf("get token error: %v", err),
		})
	}

	return c.JSON(http.StatusOK, tokenResponse)
}

func (ws *webservice) CreateUser(c echo.Context) error {
	var userDTO model.UserDTO

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user data",
		})
	}

	if userDTO.Name == "" || userDTO.Password == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: "Name and password cannot be empty",
		})
	}

	if ws.userRepository.Exists(userDTO.Name) {
		return c.JSON(http.StatusConflict, model.Response{
			Message: "User with this name already exists",
		})
	}

	id, err := ws.userRepository.CreateUser(userDTO.Name, userDTO.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, model.Response{
		Message: "Created user",
		Data:    id,
	})
}

func (ws *webservice) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid user ID",
		})
	}

	err = ws.userRepository.DeleteUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, model.Response{
		Message: "Deleted user",
	})
}
