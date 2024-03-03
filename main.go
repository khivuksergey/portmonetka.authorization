package main

import (
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"os"
)

func main() {
	cfg := webserver.LoadConfiguration("config.json")
	logger := logger.NewConsoleLogger()
	webservice := NewWebService(logger)
	router := NewRouter(&cfg.WebServer.HttpHandler, webservice)
	options := webserver.ServerOptions{
		UseLogger: true,
	}
	server := NewServer(&cfg.WebServer, router, logger, options)

	quit := make(chan os.Signal, 1)
	err := webserver.RunServer(server, &quit)
	if err != nil {
		panic(err)
	}
}
