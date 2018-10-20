package server

import (
	"github.com/gorilla/mux"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
)

//Server holds contract for all server implementation
type Server interface {
	Run(env string) error
}

type httpServer struct {
	serviceResource *serviceresource.Resource
	router          *mux.Router
	port            int
}

type execServer struct {
	serviceResource *serviceresource.Resource
	osArgs          []string
}
