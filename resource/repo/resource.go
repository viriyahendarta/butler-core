package repo

import "github.com/viriyahendarta/butler-core/infra/database"

//Resource holds resources needed for repo
type Resource struct {
	CoreDB *database.DB
}
