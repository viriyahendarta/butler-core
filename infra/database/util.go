package database

import (
	"database/sql"
	"errors"

	"github.com/viriyahendarta/butler-core/infra/errorx"
)

//OpenConnections connects to databases specified in parameter
func OpenConnections(driver string, urls ...string) ([]*sql.DB, error) {
	if len(urls) < 1 {
		return nil, errorx.New(nil, errorx.CodeDatabaseGeneral, "Failed to open database connection", errors.New("OpenConnections() method needs min 1 url parameter"))
	}

	dbs := make([]*sql.DB, 0)
	for _, url := range urls {
		db, err := sql.Open(driver, url)
		if err != nil {
			return nil, errorx.New(nil, errorx.CodeDatabaseGeneral, "Failed to open database connection", err)
		}
		dbs = append(dbs, db)
	}

	return dbs, nil
}
