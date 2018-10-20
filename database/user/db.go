package user

import (
	"context"
	"sync"

	"github.com/viriyahendarta/butler-core/database"
	d "github.com/viriyahendarta/butler-core/infra/database"

	usermodel "github.com/viriyahendarta/butler-core/model/user"
)

//Database implement database contract and holds user database implementation
type Database interface {
	database.Database

	Find(ctx context.Context, id int64) (*usermodel.User, error)
}

type userDatabase struct {
	CoreDB *d.DB
}

var db Database
var once sync.Once

//GetDatabase returns user database implementation
func GetDatabase(coreDB *d.DB) Database {
	once.Do(func() {
		db = &userDatabase{
			CoreDB: coreDB,
		}
	})

	return db
}

//Begin
func (ud *userDatabase) Begin() {

}
