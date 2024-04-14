package service

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
)

type Manager struct {
	Authorization AuthorizationService
	User          UserService
}

//go:generate mockgen -source=service.go -destination=../../service/authorization/mock/mock_authorization.go -package=mock
type AuthorizationService interface {
	Login(dto *model.UserLoginDTO) (*common.TokenResponse, error)
}

//go:generate mockgen -source=service.go -destination=../../service/user/mock/mock_user.go -package=mock
type UserService interface {
	CreateUser(dto *model.UserCreateDTO) (*uint64, error)
	UpdateUserName(dto *model.UserUpdateNameDTO) error
	UpdateUserPassword(dto *model.UserUpdatePasswordDTO) error
	DeleteUser(userId uint64) error
}
