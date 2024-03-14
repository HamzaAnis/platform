package db

import (
	"context"

	pb "github.com/HamzaAnis/platform/gen/user"
	"github.com/HamzaAnis/platform/pkg/logger"
	"github.com/HamzaAnis/platform/src/user/models"
	"github.com/jmoiron/sqlx"
)

var (
	log = logger.Logger(pb.User_ServiceDesc.ServiceName)
)

type UserDB interface {
	GetUser(context.Context, int64) (*models.User, error)
}

type userDBImpl struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) UserDB {
	return &userDBImpl{
		db: db,
	}
}

// Assert that *userDBImpl satisfies the UserDB interface
var _ UserDB = &userDBImpl{}
