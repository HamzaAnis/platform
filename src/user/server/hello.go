package server

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/user"
)

func (u *UserServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Info("Running the user app", pb.User_ServiceDesc.ServiceName)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
