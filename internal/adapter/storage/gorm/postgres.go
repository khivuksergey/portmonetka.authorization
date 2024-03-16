package gorm

import (
	"fmt"
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/env"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/entity"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm/repo"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/port/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbManager struct {
	db  *gorm.DB
	cfg *config.DBConfig
}

func NewDbManager(config config.DBConfig) (storage.IDB, error) {
	dbm := dbManager{}
	err := dbm.InitDB(config)
	if err != nil {
		return nil, err
	}
	return &dbm, err
}

func (m *dbManager) InitDB(config config.DBConfig) (err error) {
	dsn := fmt.Sprintf(config.DSN, env.PgUser, env.PgPassword, env.PgDbName, env.PgHost)

	m.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.TablePrefix,
		//},
	})

	if err != nil {
		return err
	}

	err = m.db.AutoMigrate(&entity.User{})

	return err
}

func (m *dbManager) InitRepository() *repository.Manager {
	return &repository.Manager{
		User: repo.NewUserRepository(m.db),
	}
}

func (m *dbManager) Close() (err error) {
	db, err := m.db.DB()
	if err != nil {
		return
	}
	err = db.Close()
	return
}
