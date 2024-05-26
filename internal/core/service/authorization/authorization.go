package authorization

import (
	"github.com/golang-jwt/jwt/v5"
	serviceerror "github.com/khivuksergey/portmonetka.authorization/error"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"github.com/khivuksergey/portmonetka.authorization/utility"
	"github.com/spf13/viper"
	"time"
)

type authorization struct {
	userRepository repository.UserRepository
}

func NewAuthorizationService(repositoryManager *repository.Manager) service.AuthorizationService {
	return &authorization{userRepository: repositoryManager.User}
}

func (a *authorization) Login(userLoginDTO *model.UserLoginDTO) (*model.TokenResponse, error) {
	if err := a.validateUser(userLoginDTO); err != nil {
		return nil, err
	}

	tokenResponse, err := a.getToken(*userLoginDTO)
	if err != nil {
		return nil, serviceerror.GetTokenError(err)
	}

	if err = a.userRepository.UpdateLastLoginTime(userLoginDTO.Id, tokenResponse.IssuedAt); err != nil {
		return nil, err
	}

	return tokenResponse, nil
}

func (a *authorization) validateUser(userLoginDTO *model.UserLoginDTO) error {
	if userLoginDTO.Name == "" || userLoginDTO.Password == "" {
		return serviceerror.EmptyNamePassword
	}

	user, err := a.userRepository.FindUserByName(userLoginDTO.Name)
	if err != nil {
		return err // common.UserNotFound
	}
	userLoginDTO.Id = user.Id

	if !utility.VerifyPassword(userLoginDTO.Password, user.Password) {
		return serviceerror.InvalidPassword
	}

	return nil
}

func (a *authorization) getToken(user model.UserLoginDTO) (*model.TokenResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, serviceerror.TokenClaimsFail
	}

	now := time.Now()
	expiry := now.Add(time.Hour * 24)
	if user.RememberMe {
		expiry = now.Add(time.Hour * 24 * 7)
	}

	claims["iss"] = viper.GetString("JWT_ISSUER")
	claims["sub"] = user.Id
	claims["iat"] = now.Unix()
	claims["exp"] = expiry.Unix()

	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return nil, serviceerror.TokenSignFail
	}

	return &model.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(expiry.Sub(now).Seconds()),
		IssuedAt:    now,
		Issuer:      "portmonetka.authorization",
		Subject:     user.Id,
	}, nil
}
