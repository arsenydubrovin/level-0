package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*orderRepository, error) {
	db, err := openPostgres(dsn, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	return &orderRepository{
		db: db,
	}, nil
}

func openPostgres(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return db, nil
}
