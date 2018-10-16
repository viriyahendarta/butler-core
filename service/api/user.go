package api

import (
	"net/http"
	"sync"

	"github.com/viriyahendarta/butler-core/business/user"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
)

type User interface {
	GetUserProfile(r *http.Request) (interface{}, int, error)
}

type userAPI struct {
	ServiceResource *serviceresource.Resource
}

var uAPI User
var once sync.Once

func GetUser(resource *serviceresource.Resource) User {
	once.Do(func() {
		uAPI = &userAPI{
			ServiceResource: resource,
		}
	})
	return uAPI
}

func (u *userAPI) GetUserProfile(r *http.Request) (interface{}, int, error) {
	profile, err := user.GetGetProfileBusiness(u.ServiceResource.BusinessResource).HandleBusiness(r.Context(), 1)
	if err != nil {
		return nil, -1, err
	}

	return profile, http.StatusOK, nil
}
