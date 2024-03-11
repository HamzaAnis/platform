package main

import (
	"context"
	"net"
	"os"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/pkg/postgres"

	"google.golang.org/grpc"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

type serverA struct {
	pb.UnimplementedUserServer
}

func (s *serverA) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Info("Running the user app", pb.User_ServiceDesc.ServiceName)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	postgres.NewPostgres()
	lis, err := net.Listen("tcp", ":"+os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServer(s, &serverA{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
