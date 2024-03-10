package db

import (
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func NewDb(config config.DBConfig) (db *gorm.DB) {
	user, password, dbname, host, tz := getDsnData()
	dsn := fmt.Sprintf(config.DSN, user, password, dbname, host, tz)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "portmonetka.",
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

func getDsnData() (user, password, dbname, host, timezone string) {
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB_NAME")
	host = os.Getenv("POSTGRES_HOST")
	timezone = os.Getenv("POSTGRES_TIMEZONE")
	return
}
