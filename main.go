package main

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	_ "github.com/khivuksergey/portmonetka.authorization/docs"
	"github.com/khivuksergey/portmonetka.authorization/internal/adapter/storage/gorm"
	"github.com/khivuksergey/portmonetka.authorization/internal/core/service"
	"github.com/khivuksergey/portmonetka.authorization/internal/http"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"os"
)

// @title Portmonetka authorization & user service
// @description JWT authorization and authentication. User service.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	cfg := config.LoadConfiguration("config.json")
	if err := config.LoadEnv(); err != nil {
		panic(err)
	}

	db, err := gorm.NewDbManager(cfg.DB)
	if err != nil {
		panic(err)
	}

	repositoryManager := db.InitRepository()

	services := service.NewServiceManager(repositoryManager)

	consoleLogger := logger.NewConsoleLogger()

	router := http.NewRouter(&cfg.WebServer.Router, services, consoleLogger)

	server := http.NewServer(cfg, router)

	server.AddLogger(consoleLogger)

	server.AddStopHandlers(&[]webserver.StopHandler{
		{
			Description: "Postgres",
			Stop:        db.Close,
		},
	})

	quit := make(chan os.Signal, 1)
	if err = webserver.RunServer(server, &quit); err != nil {
		panic(err)
	}
}
