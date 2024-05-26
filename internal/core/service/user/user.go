package user

import (
	serviceerror "github.com/khivuksergey/portmonetka.authorization/error"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
)

type user struct {
	userRepository repository.UserRepository
}

func NewUserService(repositoryManager *repository.Manager) service.UserService {
	return &user{userRepository: repositoryManager.User}
}

func (u *user) CreateUser(userCreateDTO model.UserCreateDTO) (uint64, error) {
	if u.userRepository.Exists(userCreateDTO.Name) {
		return 0, serviceerror.UserAlreadyExists
	}
	return u.userRepository.CreateUser(userCreateDTO.Name, userCreateDTO.Password)
}

func (u *user) DeleteUser(userId uint64) error {
	return u.userRepository.DeleteUser(userId)
}

func (u *user) UpdateUserName(userUpdateNameDTO model.UserUpdateNameDTO) error {
	if u.userRepository.Exists(userUpdateNameDTO.Name) {
		return serviceerror.UserAlreadyExists
	}
	return u.userRepository.UpdateUserName(userUpdateNameDTO.Id, userUpdateNameDTO.Name)
}

func (u *user) UpdateUserPassword(userUpdatePasswordDTO model.UserUpdatePasswordDTO) error {
	return u.userRepository.UpdateUserPassword(userUpdatePasswordDTO.Id, userUpdatePasswordDTO.Password)
}
