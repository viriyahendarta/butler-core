package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/viriyahendarta/butler-core/infra/errorx"
	usermodel "github.com/viriyahendarta/butler-core/model/user"
)

func (ud *userDatabase) Find(ctx context.Context, id int64) (*usermodel.User, error) {
	user := new(usermodel.User)

	db := ud.CoreDB.Slave()
	if err := db.Get(user, queryGetUserByID, id); err != nil {
		if err != sql.ErrNoRows {
			return nil, errorx.New(ctx, errorx.CodeQueryUser, fmt.Sprintf("Failed to find user by id: %v", id), err)
		} else {
			return nil, nil
		}
	}

	return user, nil
}
