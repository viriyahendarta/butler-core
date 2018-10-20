package service

import (
	"github.com/gorilla/mux"
	businessresource "github.com/viriyahendarta/butler-core/resource/business"
)

//Resource holds resources needed for service
type Resource struct {
	BusinessResource *businessresource.Resource
	Router           *mux.Router
}
