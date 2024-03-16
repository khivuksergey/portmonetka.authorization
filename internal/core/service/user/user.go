package user

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
)

type user struct {
	userRepository repository.UserRepository
}

func NewUserService(repo *repository.Manager) service.UserService {
	return &user{userRepository: repo.User}
}

func (u *user) CreateUser(userCreateDTO *model.UserCreateDTO) (userId *uint64, err error) {
	if err = u.validateUserCreate(userCreateDTO); err != nil {
		return
	}
	userId, err = u.userRepository.CreateUser(userCreateDTO.Name, userCreateDTO.Password)
	return
}

func (u *user) DeleteUser(userId uint64) error {
	return u.userRepository.DeleteUser(userId)
}

func (u *user) UpdateUserName(userUpdateNameDTO *model.UserUpdateNameDTO) error {
	if err := u.validateUserUpdateName(userUpdateNameDTO); err != nil {
		return err
	}
	return u.userRepository.UpdateUserName(userUpdateNameDTO.Id, userUpdateNameDTO.Name)
}

func (u *user) UpdateUserPassword(userUpdatePasswordDTO *model.UserUpdatePasswordDTO) error {
	if err := u.validateUserUpdatePassword(userUpdatePasswordDTO); err != nil {
		return err
	}
	return u.userRepository.UpdateUserPassword(userUpdatePasswordDTO.Id, userUpdatePasswordDTO.Password)
}

func (u *user) validateUserCreate(userCreateDTO *model.UserCreateDTO) error {
	if userCreateDTO.Name == "" || userCreateDTO.Password == "" {
		return common.EmptyNamePassword
	}
	// ignores deleted, but there's a unique constraint in db
	if u.userRepository.Exists(userCreateDTO.Name) {
		return common.UserAlreadyExists
	}
	return nil
}

func (u *user) validateUserUpdateName(userUpdateNameDTO *model.UserUpdateNameDTO) error {
	if userUpdateNameDTO.Name == "" {
		return common.EmptyName
	}
	if u.userRepository.Exists(userUpdateNameDTO.Name) {
		return common.UserAlreadyExists
	}
	return nil
}

func (u *user) validateUserUpdatePassword(userUpdatePasswordDTO *model.UserUpdatePasswordDTO) error {
	if userUpdatePasswordDTO.Password == "" {
		return common.EmptyPassword
	}
	return nil
}
