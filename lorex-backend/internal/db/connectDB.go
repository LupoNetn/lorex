package db

import (
	"context"
    "log/slog"
	"errors"
	"time"
	
	"github.com/luponetn/lorex/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg *config.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DBUrl)
	if err != nil {
		slog.Error("failed to parse db config", "error", err)
		return nil, errors.New("failed to parse database configuration")
	}

	// Pool tuning
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		slog.Error("failed to create connection pool", "error", err)
		return nil, errors.New("failed to create database pool")
	}

	slog.Info("database connection pool established")

	return pool, nil
}