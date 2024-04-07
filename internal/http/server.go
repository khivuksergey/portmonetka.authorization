package http

import (
	"github.com/khivuksergey/portmonetka.authorization/config"
	"github.com/khivuksergey/webserver"
	"net/http"
)

type srv struct {
	webserver.Server
}

func NewServer(
	config *config.Configuration,
	router http.Handler,
) webserver.Server {
	server := &srv{
		Server: webserver.NewServer(&config.WebServer, router, nil),
	}
	return server
}
