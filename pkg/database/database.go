package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type Config struct {
	Host    string
	Port    string
	User    string
	Passwd  string
	DBName  string
	SSLMode string
}

func NewDB(cfg Config, logger *slog.Logger) (*sql.DB, error) {
	connstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.User, cfg.Passwd, cfg.DBName, cfg.SSLMode)
	//connstr := "user=test password=test dbname=test sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
