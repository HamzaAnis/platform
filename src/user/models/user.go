package models

import (
	"time"
)

type User struct {
	UserID       int64     `db:"user_id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Balance      float64   `db:"balance"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
