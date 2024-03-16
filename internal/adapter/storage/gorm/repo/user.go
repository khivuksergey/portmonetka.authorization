package repo

import (
	"errors"
	"fmt"
	autherrors "github.com/khivuksergey/portmonetka.authorization/common"
	"github.com/khivuksergey/portmonetka.authorization/common/utility"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/entity"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/model"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db, tableName: entity.User{}.TableName()}
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

func (r *userRepository) CreateUser(name, password string) (userId *uint64, err error) {
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
	userId = &user.Id
	return
}

func (r *userRepository) UpdateUserName(id uint64, name string) error {
	return r.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("name", name).
		Error
}

func (r *userRepository) UpdateUserPassword(id uint64, password string) error {
	hashedPassword, err := utility.HashPassword(password)
	if err != nil {
		return err
	}
	return r.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("password", hashedPassword).
		Error
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
