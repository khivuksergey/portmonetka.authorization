package authorization

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/common/utility"
	"github.com/khivuksergey/portmonetka.authorization/env"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"time"
)

type authorization struct {
	userRepository repository.UserRepository
}

func NewAuthorizationService(repo *repository.Manager) service.AuthorizationService {
	return &authorization{userRepository: repo.User}
}

func (a *authorization) Login(userLoginDTO *model.UserLoginDTO) (*common.TokenResponse, error) {
	if err := a.validateUser(userLoginDTO); err != nil {
		return nil, err
	}

	tokenResponse, err := a.getToken(userLoginDTO)
	if err != nil {
		return nil, common.GetTokenError(err)
	}

	a.userRepository.UpdateLastLoginTime(userLoginDTO.Id, tokenResponse.IssuedAt)

	return tokenResponse, nil
}

func (a *authorization) validateUser(userLoginDTO *model.UserLoginDTO) error {
	if userLoginDTO.Name == "" || userLoginDTO.Password == "" {
		return common.EmptyNamePassword
	}

	user, err := a.userRepository.FindUserByName(userLoginDTO.Name)
	if err != nil {
		return err
	}
	userLoginDTO.Id = user.Id

	if !utility.VerifyPassword(userLoginDTO.Password, user.Password) {
		return common.InvalidPassword
	}

	return nil
}

func (a *authorization) getToken(user *model.UserLoginDTO) (*common.TokenResponse, error) {
	if user == nil {
		return nil, common.NilUserToken
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, common.TokenClaimsFail
	}

	now := time.Now()
	expiry := now.Add(time.Hour * 24)
	if user.RememberMe {
		expiry = now.Add(time.Hour * 24 * 7)
	}

	claims["iss"] = env.JwtIssuer
	claims["sub"] = user.Id
	claims["iat"] = now.Unix()
	claims["exp"] = expiry.Unix()

	tokenString, err := token.SignedString(env.JwtSecret)
	if err != nil {
		return nil, err
	}

	return &common.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(expiry.Sub(now).Seconds()),
		IssuedAt:    now,
		Issuer:      "portmonetka.authorization",
		Subject:     user.Id,
	}, nil
}
