package db

import (
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDb(config config.DBConfig) (db *gorm.DB) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf(config.DSN, os.Getenv("POSTGRES_PASSWORD")),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return
}
