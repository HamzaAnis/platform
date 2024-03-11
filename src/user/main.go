package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HamzaAnis/platform/gen/user"

	"google.golang.org/grpc"
)

type serverA struct {
	pb.UnimplementedUserServer
	// Add NATS connection if needed
}

func (s *serverA) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Logic to interact with ServiceB over NATS or directly return a response
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &serverA{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
