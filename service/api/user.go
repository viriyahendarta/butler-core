package api

import (
	"net/http"
	"sync"

	"github.com/viriyahendarta/butler-core/infra/contextx"

	"github.com/viriyahendarta/butler-core/business/user"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
)

//User holds user services implementation
type User interface {
	GetUserInfo(r *http.Request) (interface{}, int, error)
}

type userAPI struct {
	ServiceResource *serviceresource.Resource
}

var uAPI User
var once sync.Once

//GetUser returns user service
func GetUser(resource *serviceresource.Resource) User {
	once.Do(func() {
		uAPI = &userAPI{
			ServiceResource: resource,
		}
	})
	return uAPI
}

//GetUserInfo handles Get User Info request
func (u *userAPI) GetUserInfo(r *http.Request) (interface{}, int, error) {
	authID, err := contextx.GetAuthID(r.Context())
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}

	info, err := user.GetGetInfoBusiness(u.ServiceResource.BusinessResource).HandleBusiness(r.Context(), authID)
	if err != nil {
		return nil, -1, err
	}

	return info, http.StatusOK, nil
}
