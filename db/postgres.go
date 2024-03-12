package db

import (
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/env"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDb(config config.DBConfig) (db *gorm.DB) {
	dsn := fmt.Sprintf(config.DSN, env.PgUser, env.PgPassword, env.PgDbName, env.PgHost, env.PgTimezone)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.TablePrefix,
		},
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return
}
