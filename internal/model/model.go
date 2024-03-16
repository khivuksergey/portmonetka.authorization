package model

import (
	"gorm.io/gorm"
	"time"
)

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
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserUpdatePasswordDTO struct {
	Id       uint64 `json:"id"`
	Password string `json:"password"`
}
