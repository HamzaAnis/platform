package jwt

import (
	"context"

	"github.com/HamzaAnis/platform/pkg/logger"
)

var (
	log = logger.Logger("JWT")
)

type JWT interface {
	GenerateToken(int64) (*string, error)
	VerifyToken(context.Context, string) (*CustomClaims, error)
}

type jwtImpl struct {
	secret string
}

func NewJWT(secret string) JWT {
	return &jwtImpl{
		secret: secret,
	}
}

// Assert that *userServiceImpl satisfies the UserServiceServer interface
var _ JWT = &jwtImpl{}
