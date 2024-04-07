package model

type UserCreateDTO struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginDTO struct {
	Id         uint64 `json:"-"`
	Name       string `json:"name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
}

type UserUpdateNameDTO struct {
	Id   uint64 `json:"-"`
	Name string `json:"name" validate:"required"`
}

type UserUpdatePasswordDTO struct {
	Id       uint64 `json:"-"`
	Password string `json:"password" validate:"required"`
}
