package api

import (
	"github.com/HamzaAnis/platform/pkg/postgres"
	"github.com/HamzaAnis/platform/src/transaction/db"
)

type TransactionAPI interface {
}

type transactionAPIImpl struct {
	db db.TransactionDB
}

func NewTransactionAPI() TransactionAPI {
	return &transactionAPIImpl{
		db: db.NewTransactionDB(postgres.NewPostgres().GetDB()),
	}
}

// Assert that *transactionAPIImpl satisfies the TransactionAPI interface
var _ TransactionAPI = &transactionAPIImpl{}
