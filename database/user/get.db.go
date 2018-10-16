package user

import (
	"context"
	"database/sql"
	"fmt"

	e "github.com/viriyahendarta/butler-core/infra/error"
	usermodel "github.com/viriyahendarta/butler-core/model/user"
)

func (ud *userDatabase) Find(ctx context.Context, id int64) (*usermodel.User, error) {
	user := new(usermodel.User)

	db := ud.CoreDB.Slave()
	if err := db.Get(user, queryGetUserByID, id); err != nil {
		if err != sql.ErrNoRows {
			return nil, e.New(ctx, e.CodeQueryUser, fmt.Sprintf("Failed to find user by id: %v", id), err)
		} else {
			return nil, nil
		}
	}

	return user, nil
}
