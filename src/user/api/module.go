package api

import (
	"github.com/HamzaAnis/platform/pkg/postgres"
	"github.com/HamzaAnis/platform/src/user/db"
)

type UserAPI interface {
}

type userAPIImpl struct {
	db db.UserDB
}

func NewUserAPI() UserAPI {
	return &userAPIImpl{
		db: db.NewUserDB(postgres.NewPostgres().GetDB()),
	}
}

// Assert that *userAPIImpl satisfies the UserAPI interface
var _ UserAPI = &userAPIImpl{}
