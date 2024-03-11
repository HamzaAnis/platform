package server

import (
	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/user/api"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

type UserServerImpl struct {
	pb.UserServer
	api api.UserAPI
}

func NewUserServer() pb.UserServer {
	api := api.NewUserAPI()
	return &UserServerImpl{
		api: api,
	}
}

// Assert that *userServiceImpl satisfies the UserServiceServer interface
var _ pb.UserServer = &UserServerImpl{}
