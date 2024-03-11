package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

func (p *postgresImpl) GetDB() *sql.DB {
	return p.db
}

func connect() (*sql.DB, error) {
	fmt.Println(os.Getenv("POSTGRES_PORT"))
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Info("⛁ Connected to MySQL Database!")
	return db, nil
}
