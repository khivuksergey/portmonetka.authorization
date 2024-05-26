package service

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
)

type Manager struct {
	Authorization AuthorizationService
	User          UserService
}

type AuthorizationService interface {
	Login(dto *model.UserLoginDTO) (*model.TokenResponse, error)
}

type UserService interface {
	CreateUser(dto model.UserCreateDTO) (uint64, error)
	UpdateUserName(dto model.UserUpdateNameDTO) error
	UpdateUserPassword(dto model.UserUpdatePasswordDTO) error
	DeleteUser(userId uint64) error
}
