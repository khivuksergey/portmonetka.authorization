package entity

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

func (User) TableName() string { return "portmonetka.users" }
