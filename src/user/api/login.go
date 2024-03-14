package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *userAPIImpl) Login(ctx context.Context, userID int64) (*string, error) {
	user, err := u.db.GetUser(ctx, userID)
	if user == nil || err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return u.jwt.GenerateToken(userID)
}
