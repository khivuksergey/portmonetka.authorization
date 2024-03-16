package common

import "time"

type Response struct {
	Message string
	Data    any
}

type TokenResponse struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int64     `json:"expires_in"`
	IssuedAt    time.Time `json:"issued_at"`
	Issuer      string    `json:"issuer"`
	Subject     uint64    `json:"subject"`
}
