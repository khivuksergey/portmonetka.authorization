package repository

import (
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"time"
)

type Manager struct {
	User UserRepository
}

type UserRepository interface {
	Exists(name string) bool
	FindUserByName(name string) (*model.User, error)
	CreateUser(name, password string) (*uint64, error)
	UpdateUserName(id uint64, name string) error
	UpdateUserPassword(id uint64, password string) error
	DeleteUser(id uint64) error
	UpdateLastLoginTime(userId uint64, loginTime time.Time)
}
