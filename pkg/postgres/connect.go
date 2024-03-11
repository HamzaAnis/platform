package postgres

import (
	"database/sql"
	"fmt"

	"github.com/HamzaAnis/platform/pkg/config"
)

func (p *postgresImpl) GetDB() *sql.DB {
	return p.db
}

func connect() (*sql.DB, error) {
	cfg := config.Cfg
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info("⛁ Connected to MySQL Database!")
	return db, nil
}
