package server

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/transaction"
)

func (u *transactionServerImpl) Up(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionReply, error) {
	log.Infof("Running %v", pb.Transaction_ServiceDesc.ServiceName)
	return &pb.TransactionReply{Message: "ok"}, nil
}
