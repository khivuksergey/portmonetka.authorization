package db

import (
	"errors"
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	autherrors "github.com/khivuksergey/portmonetka.authorization/errors"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"github.com/khivuksergey/portmonetka.authorization/utility"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Exists(name string) bool
	FindUserByName(name string) (*model.User, error)
	FindUserById(id uint64) (*model.User, error)
	CreateUser(name, password string) (uint64, error)
	UpdateUser(user model.User) error
	DeleteUser(id uint64) error
	UpdateLastLoginTime(userId uint64, loginTime time.Time)
}

type userRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(dbConfig config.DBConfig) UserRepository {
	return &userRepository{
		NewDb(dbConfig),
		dbConfig.TablePrefix + "users",
	}
}

func (r *userRepository) Exists(name string) bool {
	var count int64
	r.db.Model(&model.User{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (r *userRepository) FindUserByName(name string) (user *model.User, err error) {
	result := r.db.Where("name = ?", name).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = autherrors.UserNotFound
	}
	return
}

func (r *userRepository) FindUserById(id uint64) (user *model.User, err error) {
	result := r.db.Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = autherrors.UserNotFound
	}
	return
}

func (r *userRepository) CreateUser(name, password string) (userId uint64, err error) {
	hashedPassword, err := utility.HashPassword(password)
	if err != nil {
		return
	}
	user := &model.User{
		Name:     name,
		Password: hashedPassword,
	}
	if err = r.db.Create(user).Error; err != nil {
		return
	}
	userId = user.Id
	return
}

func (r *userRepository) UpdateUser(user model.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id uint64) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) UpdateLastLoginTime(userId uint64, loginTime time.Time) {
	query := fmt.Sprintf("UPDATE %s SET last_login = ? WHERE id = ?", r.tableName)
	if err := r.db.Exec(query, loginTime, userId).Error; err != nil {
		fmt.Printf("Error updating last login time: %s. UserId: %d, login time: %v\n",
			err.Error(), userId, loginTime)
	}
}
