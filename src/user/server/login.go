package server

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/user"
)

func (u *UserServerImpl) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := u.api.Login(ctx, in.GetUserID())
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: *token}, nil
}
