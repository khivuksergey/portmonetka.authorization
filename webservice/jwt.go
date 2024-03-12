package webservice

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/khivuksergey/portmonetka.authorization/env"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"time"
)

func (ws *webservice) getToken(user *model.UserLoginDTO) (*model.TokenResponse, error) {
	if user == nil {
		return nil, errors.New("cannot get token for nil user")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to get token claims")
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

	return &model.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   int64(expiry.Sub(now).Seconds()),
		IssuedAt:    now,
		Issuer:      "portmonetka.authorization",
		Subject:     user.Id,
	}, nil
}
