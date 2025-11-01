package storage

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func init() {
	_ = godotenv.Load("/root/.env")
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("Database URL is required")
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database: %v", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Unable to connect with the Database: %v", err)
	}

	Pool = pool
	log.Println("Connected to the database")
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
