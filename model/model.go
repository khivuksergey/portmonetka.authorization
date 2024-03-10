package model

import "gorm.io/gorm"

type Response struct {
	Message string
	Data    any
}

type UserDTO struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u User) MapToDTO() *UserDTO {
	return &UserDTO{
		Name:     u.Name,
		Password: u.Password,
	}
}
