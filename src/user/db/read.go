package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/HamzaAnis/platform/src/user/models"
)

func (d *userDBImpl) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	if err := d.db.Get(&user, `SELECT user_id from users where user_id = $1`, userID); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return &user, nil
}
