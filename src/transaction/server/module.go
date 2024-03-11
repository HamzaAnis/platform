package server

import (
	pb "github.com/HamzaAnis/platform/gen/transaction"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/transaction/api"
)

var (
	log = logger.Logger(pb.Transaction_ServiceDesc.ServiceName)
)

type transactionServerImpl struct {
	pb.TransactionServer
	api api.TransactionAPI
}

func NewTransactionServer() pb.TransactionServer {
	api := api.NewTransactionAPI()
	return &transactionServerImpl{
		api: api,
	}
}

// Assert that *userServiceImpl satisfies the UserServiceServer interface
var _ pb.TransactionServer = &transactionServerImpl{}
