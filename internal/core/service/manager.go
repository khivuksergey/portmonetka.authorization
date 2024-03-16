package service

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/authorization"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/user"
)

func NewServiceManager(repo *repository.Manager) *service.Manager {
	return &service.Manager{
		Authorization: authorization.NewAuthorizationService(repo),
		User:          user.NewUserService(repo),
	}
}
