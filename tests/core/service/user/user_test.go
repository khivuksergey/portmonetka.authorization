package user

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm/repo/mock"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/user"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreateUser_Success(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userCreateDTO := &model.UserCreateDTO{
		Name:     "new_user",
		Password: "password",
	}

	userCreateId := uint64(1)

	mockUserRepository.
		EXPECT().
		Exists(userCreateDTO.Name).
		Times(1).
		Return(false)

	mockUserRepository.
		EXPECT().
		CreateUser(userCreateDTO.Name, userCreateDTO.Password).
		Times(1).
		Return(&userCreateId, nil)

	createdUserId, err := userService.CreateUser(userCreateDTO)

	assert.NoError(t, err)
	assert.NotNil(t, createdUserId)
	assert.Equal(t, *createdUserId, userCreateId)
}

func TestCreateUser_AlreadyExists_Error(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userCreateDTO := &model.UserCreateDTO{
		Name:     "existing_user",
		Password: "password",
	}

	mockUserRepository.
		EXPECT().
		Exists(userCreateDTO.Name).
		Times(1).
		Return(true)

	createdUserId, err := userService.CreateUser(userCreateDTO)

	assert.ErrorIs(t, err, common.UserAlreadyExists)
	assert.Nil(t, createdUserId)
}

func TestDeleteUser_Success(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userDeleteId := uint64(1)

	mockUserRepository.
		EXPECT().
		DeleteUser(userDeleteId).
		Times(1).
		Return(nil)

	err := userService.DeleteUser(userDeleteId)

	assert.NoError(t, err)
}

func TestUpdateUserName_Success(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userUpdateNameDTO := &model.UserUpdateNameDTO{
		Id:   1,
		Name: "new_unique_username",
	}

	mockUserRepository.
		EXPECT().
		Exists(userUpdateNameDTO.Name).
		Times(1).
		Return(false)

	mockUserRepository.
		EXPECT().
		UpdateUserName(userUpdateNameDTO.Id, userUpdateNameDTO.Name).
		Times(1).
		Return(nil)

	err := userService.UpdateUserName(userUpdateNameDTO)

	assert.NoError(t, err)
}

func TestUpdateUserName_NameAlreadyExists_Error(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userUpdateNameDTO := &model.UserUpdateNameDTO{
		Id:   1,
		Name: "existing_username",
	}

	mockUserRepository.
		EXPECT().
		Exists(userUpdateNameDTO.Name).
		Times(1).
		Return(true)

	err := userService.UpdateUserName(userUpdateNameDTO)

	assert.ErrorIs(t, err, common.UserAlreadyExists)
}

func TestUpdateUserPassword_Success(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	userService := user.NewUserService(mockManager)

	userUpdatePasswordDTO := &model.UserUpdatePasswordDTO{
		Id:       1,
		Password: "new_password",
	}

	mockUserRepository.
		EXPECT().
		UpdateUserPassword(userUpdatePasswordDTO.Id, userUpdatePasswordDTO.Password).
		Times(1).
		Return(nil)

	err := userService.UpdateUserPassword(userUpdatePasswordDTO)

	assert.NoError(t, err)
}
