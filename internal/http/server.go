package http

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
)

const configPath = "config.json"

func NewServer() webserver.Server {
	config.LoadEnv()

	cfg := config.LoadConfiguration(configPath)

	db := gorm.NewDbManager(cfg.DB)

	services := service.NewServiceManager(db.InitRepositoryManager())

	log := logger.Default.SetLevel(logger.GetLogLevelFromString(cfg.Logger.LogLevel))

	router := NewRouter(cfg, services, log)

	server := webserver.NewServer(router).
		AddLogger(log).
		AddStopHandlers(webserver.NewStopHandler("Postgres", db.Close))

	return server
}
