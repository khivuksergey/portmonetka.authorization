package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"time"
)

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int64     `json:"expires_in"`
	IssuedAt    time.Time `json:"issued_at"`
	Issuer      string    `json:"issuer"`
	Subject     string    `json:"subject"`
}

func (ws *webservice) getToken(user *model.UserDTO) (*TokenResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	if user.RememberMe {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	response := &TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   claims["exp"].(int64) - time.Now().Unix(),
		IssuedAt:    time.Now(),
		Issuer:      "portmonetka.authorization",
		Subject:     user.Name,
	}

	return response, nil
}
