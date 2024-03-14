package api

import (
	"context"
	"os"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/jwt"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/pkg/postgres"
	"github.com/HamzaAnis/platform/src/user/db"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

type UserAPI interface {
	Login(ctx context.Context, userID int64) (*string, error)
	GetBalanceAPI(context.Context) (float64, error)
}

type userAPIImpl struct {
	db  db.UserDB
	jwt jwt.JWT
}

func NewUserAPI() UserAPI {
	return &userAPIImpl{
		db:  db.NewUserDB(postgres.NewPostgres().GetDB()),
		jwt: jwt.NewJWT(os.Getenv("JWT_SECRET")),
	}
}

// Assert that *userAPIImpl satisfies the UserAPI interface
var _ UserAPI = &userAPIImpl{}
