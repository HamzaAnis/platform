package db

import "database/sql"

type TransactionDB interface {
}

type transactionDBImpl struct {
	db *sql.DB
}

func NewTransactionDB(db *sql.DB) TransactionDB {
	return &transactionDBImpl{
		db: db,
	}
}

// Assert that *userDBImpl satisfies the UserDB interface
var _ TransactionDB = &transactionDBImpl{}
