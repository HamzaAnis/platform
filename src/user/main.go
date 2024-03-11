package main

import (
	"context"
	"net"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/config"
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
	err := config.LoadConfig("docker/env/user-service.env")
	if err != nil {
		log.Fatal(err)
	}

	postgres.NewPostgres()
	lis, err := net.Listen("tcp", ":"+config.Cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServer(s, &serverA{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
