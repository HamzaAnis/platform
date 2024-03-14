package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/HamzaAnis/platform/src/user/models"
)

func (d *userDBImpl) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	if err := d.db.Get(&user, `SELECT user_id, username from users where user_id = $1`, userID); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return &user, nil
}

func (d *userDBImpl) GetUserBalance(ctx context.Context) (float64, error) {
	var balance float64
	userId := ctx.Value("userID")
	if err := d.db.Get(&balance, `SELECT balance from users where user_id = $1`, userId); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			return 0, errors.New("balance not found")
		}
		return 0, err
	}
	return balance, nil
}
