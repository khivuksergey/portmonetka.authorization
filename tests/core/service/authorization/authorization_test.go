package authorization

import (
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/common/utility"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm/repo/mock"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service/authorization"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	authorizationService := authorization.NewAuthorizationService(mockManager)

	userLoginDTO := &model.UserLoginDTO{
		Name:     "valid_user",
		Password: "valid_password",
	}

	hashedValidPassword, _ := utility.HashPassword("valid_password")
	expectedUser := &model.User{Id: 1, Password: hashedValidPassword}

	mockUserRepository.
		EXPECT().
		FindUserByName(userLoginDTO.Name).
		Times(1).
		Return(expectedUser, nil)

	mockUserRepository.
		EXPECT().
		UpdateLastLoginTime(expectedUser.Id, gomock.Any()).
		Times(1)

	token, err := authorizationService.Login(userLoginDTO)

	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestLogin_InvalidPassword_Error(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	authorizationService := authorization.NewAuthorizationService(mockManager)

	userLoginDTO := &model.UserLoginDTO{
		Name:     "valid_user",
		Password: "invalid_password",
	}

	hashedValidPassword, _ := utility.HashPassword("valid_password")
	expectedUser := &model.User{Id: 1, Password: hashedValidPassword}

	mockUserRepository.
		EXPECT().
		FindUserByName(userLoginDTO.Name).
		Times(1).
		Return(expectedUser, nil)

	token, err := authorizationService.Login(userLoginDTO)

	assert.ErrorIs(t, err, common.InvalidPassword)
	assert.Nil(t, token)
}

func TestLogin_EmptyNamePassword_Error(t *testing.T) {
	ctl := gomock.NewController(t)

	mockUserRepository := mock.NewMockUserRepository(ctl)
	mockManager := &repository.Manager{
		User: mockUserRepository,
	}

	authorizationService := authorization.NewAuthorizationService(mockManager)

	userLoginDTO := &model.UserLoginDTO{}

	token, err := authorizationService.Login(userLoginDTO)

	assert.ErrorIs(t, err, common.EmptyNamePassword)
	assert.Nil(t, token)
}
