package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	_context "github.com/viriyahendarta/butler-core/infra/context"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
	"github.com/viriyahendarta/butler-core/service/api"
)

func InitHTTPServer(serviceResource *serviceresource.Resource, port int) Server {
	return &httpServer{
		serviceResource: serviceResource,
		router:          serviceResource.Router,
		port:            port,
	}
}

func (s *httpServer) Run(env string) error {
	s.registerAPI()

	address := fmt.Sprint("0.0.0.0:", s.port)
	log.Printf("Starting [%s] server at %s\n", env, address)
	return gracehttp.Serve(&http.Server{
		Addr:         address,
		Handler:      s.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	})
}

func (s *httpServer) registerAPI() {
	router := s.router.PathPrefix("/api")

	userRouter := router.PathPrefix("/user")
	userAPI := api.GetUser(s.serviceResource)
	userRouter.Path("/profile").Methods(http.MethodGet).HandlerFunc(s.handleAPI(userAPI.GetUserProfile))
}

func (s *httpServer) handleAPI(handler api.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := _context.AppendStartTime(r.Context())
		result, successCode, err := handler(r.WithContext(ctx))
		s.serviceResource.RenderJSON(ctx, w, result, successCode, err)
	}
}
