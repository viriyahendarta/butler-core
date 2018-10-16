package user

import (
	"context"
	"fmt"
	"sync"

	"github.com/viriyahendarta/butler-core/business"
	userdb "github.com/viriyahendarta/butler-core/database/user"
	userdomain "github.com/viriyahendarta/butler-core/domain/user"
	e "github.com/viriyahendarta/butler-core/infra/error"
	businessresource "github.com/viriyahendarta/butler-core/resource/business"
)

type getProfileBusiness struct {
	userDB userdb.Database
}

type GetProfileBusiness interface {
	business.Business
	HandleBusiness(ctx context.Context, userID int64) (*userdomain.Profile, error)
}

var bGetProfile GetProfileBusiness
var once sync.Once

func GetGetProfileBusiness(resource *businessresource.Resource) GetProfileBusiness {
	once.Do(func() {
		bGetProfile = &getProfileBusiness{
			userDB: resource.UserDB,
		}
	})
	return bGetProfile
}

func (b *getProfileBusiness) HandleBusiness(ctx context.Context, userID int64) (*userdomain.Profile, error) {
	user, err := b.userDB.Find(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, e.New(ctx, e.CodeBadRequest, fmt.Sprintf("User with id [%v] is not exists", userID), nil)
	}

	return &userdomain.Profile{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
