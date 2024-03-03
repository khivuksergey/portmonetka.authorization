package main

import (
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"net/http"
)

type srv struct {
	server  webserver.Server
	service WebService
}

func NewServer(config *webserver.WebServerConfig, router http.Handler, logger logger.Logger, options webserver.ServerOptions) webserver.Server {
	srv := &srv{
		server: webserver.NewServer(config, router, logger, options),
	}
	return srv.server
}
