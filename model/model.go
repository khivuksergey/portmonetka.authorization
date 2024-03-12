package model

import (
	"gorm.io/gorm"
	"time"
)

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

type User struct {
	Id         uint64         `gorm:"primarykey"`
	Name       string         `gorm:"unique;not null"`
	Password   string         `gorm:"not null"`
	RememberMe bool           `gorm:"-"`
	CreatedAt  time.Time      `gorm:"<-:create"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UpdatedAt  time.Time
	LastLogin  time.Time
}

type UserCreateDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserLoginDTO struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type UserUpdateNameDTO struct {
	Name string `json:"name"`
}

type UserUpdatePasswordDTO struct {
	Password string `json:"password"`
}
