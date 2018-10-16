package service

import (
	"github.com/gorilla/mux"
	businessresource "github.com/viriyahendarta/butler-core/resource/business"
)

type Resource struct {
	BusinessResource *businessresource.Resource
	Router           *mux.Router
}
