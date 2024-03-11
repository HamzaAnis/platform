package main

import (
	"net"
	"os"

	pb "github.com/HamzaAnis/platform/gen/transaction"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/transaction/server"

	"google.golang.org/grpc"
)

var (
	log = logger.Logger(pb.Transaction_ServiceDesc.ServiceName)
)

type serverB struct {
	pb.TransactionServer
}

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	serverImpl := server.NewTransactionServer()

	pb.RegisterTransactionServer(s, serverImpl)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
