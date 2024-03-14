package db

import (
	"github.com/jmoiron/sqlx"
)

type TransactionDB interface {
}

type transactionDBImpl struct {
	db *sqlx.DB
}

func NewTransactionDB(db *sqlx.DB) TransactionDB {
	return &transactionDBImpl{
		db: db,
	}
}

// Assert that *userDBImpl satisfies the UserDB interface
var _ TransactionDB = &transactionDBImpl{}
