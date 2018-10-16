package business

import "github.com/viriyahendarta/butler-core/database/user"

//Resource holds resources needed for business
type Resource struct {
	UserDB user.Database
}
