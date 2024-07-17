package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	var pool *pgxpool.Pool
	var err error

	// Retry connection
	for i := 0; i < 5; i++ {
		pool, err = pgxpool.Connect(context.Background(), connStr)
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database. Retrying in 5 seconds...\n")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database after 5 attempts: %v", err)
	}

	return pool, nil
}
