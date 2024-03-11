package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"

	"google.golang.org/grpc"
)

type serverA struct {
	pb.UnimplementedUserServer
	// Add NATS connection if needed
	logger logger.Logger
}

func (s *serverA) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.logger.Log().Info("Running the user app")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServer(s, &serverA{
		logger: logger.NewLogger("user"),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
