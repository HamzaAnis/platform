package postgres

import (
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	log = logger.Logger("DB")
)

type Postgres interface {
	GetDB() *sqlx.DB
}

type postgresImpl struct {
	db *sqlx.DB
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
