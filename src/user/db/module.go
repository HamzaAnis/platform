package db

import "database/sql"

type UserDB interface {
}

type userDBImpl struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) UserDB {
	return &userDBImpl{
		db: db,
	}
}

// Assert that *userDBImpl satisfies the UserDB interface
var _ UserDB = &userDBImpl{}
