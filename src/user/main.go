package main

import (
	"net"
	"os"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/user/server"

	"google.golang.org/grpc"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("SERVICE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	serverImpl := server.NewUserServer()

	pb.RegisterUserServer(s, serverImpl)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
