package api

import (
	"context"
)

func (u *userAPIImpl) GetBalanceAPI(ctx context.Context) (float64, error) {
	claims, err := u.jwt.VerifyToken(ctx)
	if err != nil {
		return 0, err
	}
	ctx = context.WithValue(ctx, "userID", claims.UserID)
	return u.db.GetUserBalance(ctx)
}
