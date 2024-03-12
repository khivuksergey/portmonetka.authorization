package main

import (
	"github.com/khivuksergey/portmonetka.authorization/webservice"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"net/http"
)

type srv struct {
	webserver.Server
	service webservice.WebService
}

func NewServer(
	config *webserver.WebServerConfig,
	router http.Handler,
	logger logger.Logger,
	options webserver.ServerOptions,
) webserver.Server {
	server := &srv{
		Server: webserver.NewServer(config, router, logger, options),
	}
	return server
}
