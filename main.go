package main

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"os"
)

func main() {
	loadEnv()
	cfg := config.LoadConfiguration("config.json")
	consoleLogger := logger.NewConsoleLogger()
	webservice := NewWebService(cfg, consoleLogger)
	router := NewRouter(&cfg.WebServer.HttpHandler, webservice)
	options := webserver.ServerOptions{
		UseLogger: true,
	}
	server := NewServer(&cfg.WebServer, router, consoleLogger, options)

	quit := make(chan os.Signal, 1)
	err := webserver.RunServer(server, &quit)
	if err != nil {
		panic(err)
	}
}
