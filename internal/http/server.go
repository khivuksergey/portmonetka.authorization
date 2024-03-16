package http

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/webserver"
	"github.com/khivuksergey/webserver/logger"
	"net/http"
)

type srv struct {
	webserver.Server
}

func NewServer(
	config *config.Configuration,
	router http.Handler,
	logger logger.Logger,
	options webserver.ServerOptions,
) webserver.Server {
	server := &srv{
		Server: webserver.NewServer(&config.WebServer, router, logger, options),
	}
	return server
}
