package main

import (
	"errors"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/db"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"github.com/khivuksergey/portmonetka.authorization/utility"
	"gorm.io/gorm"
)

var (
	UserNotFound    = errors.New("user was not found")
	InvalidPassword = errors.New("invalid password")
)

type UserRepository interface {
	Exists(name string) bool
	FindUserByName(name string) (*model.UserDTO, error)
	CreateUser(name, password string) (uint, error)
	UpdateUser(user model.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConfig config.DBConfig) UserRepository {
	return &userRepository{
		db.NewDb(dbConfig),
	}
}

func (r *userRepository) Exists(name string) bool {
	var count int64
	r.db.Model(&model.User{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (r *userRepository) FindUserByName(name string) (*model.UserDTO, error) {
	var user model.User
	result := r.db.Where("name = ?", name).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, UserNotFound
	}
	return user.MapToDTO(), nil
}

func (r *userRepository) CreateUser(name, password string) (uint, error) {
	hashedPassword, err := utility.HashPassword(password)
	if err != nil {
		return 0, err
	}
	user := &model.User{
		Name:     name,
		Password: hashedPassword,
	}
	if err = r.db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *userRepository) UpdateUser(user model.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
