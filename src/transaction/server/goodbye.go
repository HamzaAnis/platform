package server

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/transaction"
)

func (u *transactionServerImpl) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	log.Infof("Running %v", pb.Transaction_ServiceDesc.ServiceName)
	return &pb.GoodbyeReply{Message: "Good bye " + in.Name}, nil
}
