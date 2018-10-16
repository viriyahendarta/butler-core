package database

import (
	"database/sql"
	"sync"
	"sync/atomic"

	"github.com/jmoiron/sqlx"
)

const (
	PostgresDriver string = "postgres"
)

type DB struct {
	mtx sync.RWMutex

	driver string
	dbs    []*sql.DB

	counter uint64
}

type PrepStmt string

func New(driver string, dbs []*sql.DB) *DB {
	db := &DB{
		driver: driver,
		dbs:    dbs,
	}

	return db
}

func (db *DB) Master() *sqlx.DB {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	return sqlx.NewDb(db.dbs[0], db.driver)
}

func (db *DB) Slave() *sqlx.DB {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	atomic.AddUint64(&db.counter, 1)
	slavesLen := uint64(len(db.dbs) - 1)
	if slavesLen < 1 {
		return sqlx.NewDb(db.dbs[0], db.driver)
	}

	return sqlx.NewDb(db.dbs[(db.counter%slavesLen)+1], db.driver)
}
