package database

import (
	"database/sql"
	"errors"

	"github.com/viriyahendarta/butler-core/infra/errorx"
)

func OpenConnections(driver string, urls ...string) ([]*sql.DB, error) {
	if len(urls) < 1 {
		return nil, errorx.New(nil, errorx.CodeDatabaseGeneral, "Failed to open database connection", errors.New("OpenConnections() method needs min 1 url parameter"))
	}

	dbs := make([]*sql.DB, 0)
	for _, url := range urls {
		if db, err := sql.Open(driver, url); err != nil {
			return nil, errorx.New(nil, errorx.CodeDatabaseGeneral, "Failed to open database connection", err)
		} else {
			dbs = append(dbs, db)
		}
	}

	return dbs, nil
}
