package server

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/user"
)

func (u *UserServerImpl) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceReply, error) {
	balance, err := u.api.GetBalanceAPI(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetBalanceReply{Balance: balance}, nil
}
