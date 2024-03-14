package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/user/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

func main() {
	addr := ":" + os.Getenv("SERVICE_PORT")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	serverImpl := server.NewUserServer()

	pb.RegisterUserServer(s, serverImpl)

	grpcMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// Register gRPC-Gateway with the ServeMux
	err = pb.RegisterUserHandlerFromEndpoint(context.Background(), grpcMux, addr, opts)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				grpcMux.ServeHTTP(w, r)
			}
		}), &http2.Server{}),
	}

	if err := httpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
