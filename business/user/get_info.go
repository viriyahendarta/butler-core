package user

import (
	"context"
	"fmt"
	"sync"

	"github.com/viriyahendarta/butler-core/business"
	userdb "github.com/viriyahendarta/butler-core/database/user"
	userdomain "github.com/viriyahendarta/butler-core/domain/user"
	"github.com/viriyahendarta/butler-core/infra/errorx"
	businessresource "github.com/viriyahendarta/butler-core/resource/business"
)

type getInfoBusiness struct {
	userDB userdb.Database
}

//GetInfoBusiness implement business contract for Get Info
type GetInfoBusiness interface {
	business.Business
	HandleBusiness(ctx context.Context, userID int64) (*userdomain.Info, error)
}

var bGetInfo GetInfoBusiness
var once sync.Once

//GetGetInfoBusiness returns GetInfoBusiness implementation
func GetGetInfoBusiness(resource *businessresource.Resource) GetInfoBusiness {
	once.Do(func() {
		bGetInfo = &getInfoBusiness{
			userDB: resource.UserDB,
		}
	})
	return bGetInfo
}

//HandleBusiness handles Get Info business processs
func (b *getInfoBusiness) HandleBusiness(ctx context.Context, userID int64) (*userdomain.Info, error) {
	user, err := b.userDB.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errorx.New(ctx, errorx.CodeBadRequest, fmt.Sprintf("User with id [%v] is not exists", userID), nil)
	}

	return &userdomain.Info{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
