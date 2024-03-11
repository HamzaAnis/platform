package main

import (
	"context"
	"log"
	"net"

	pb "github.com/HamzaAnis/platform/gen/transaction"

	"google.golang.org/grpc"
)

type serverB struct {
	pb.UnimplementedTransactionServer
}

func (s *serverB) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	return &pb.GoodbyeReply{Message: "Goodbye there " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTransactionServer(s, &serverB{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve1: %v", err)
	}
}
