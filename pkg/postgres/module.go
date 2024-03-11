package postgres

import (
	"database/sql"

	"github.com/HamzaAnis/platform/pkg/logger"
	_ "github.com/lib/pq"
)

var (
	log = logger.Logger("DB")
)

type Postgres interface {
	GetDB() *sql.DB
}

type postgresImpl struct {
	db *sql.DB
}

// Assert that *postgresImpl satisfies the Postgres interface
var _ Postgres = &postgresImpl{}

func NewPostgres() Postgres {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	return &postgresImpl{
		db: db,
	}
}
