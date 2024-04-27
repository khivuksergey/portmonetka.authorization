package service

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/authorization"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/user"
)

func NewServiceManager(repositoryManager *repository.Manager) *service.Manager {
	return &service.Manager{
		Authorization: authorization.NewAuthorizationService(repositoryManager),
		User:          user.NewUserService(repositoryManager),
	}
}
