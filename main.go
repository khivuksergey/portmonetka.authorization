package main

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/portmonetka.authorization/env"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/http"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"os"
)

func main() {
	env.LoadEnv()
	cfg := config.LoadConfiguration("config.json")

	consoleLogger := logger.NewConsoleLogger()

	db, err := gorm.NewDbManager(cfg.DB)
	if err != nil {
		panic(err)
	}

	repositoryManager := db.InitRepository()

	services := service.NewServiceManager(repositoryManager)

	router := http.NewRouter(&cfg.WebServer.HttpHandler, services, consoleLogger)

	options := webserver.ServerOptions{
		UseLogger: true,
	}
	server := http.NewServer(cfg, router, consoleLogger, options)

	quit := make(chan os.Signal, 1)
	if webserver.RunServer(server, &quit) != nil {
		panic(err)
	}
}
