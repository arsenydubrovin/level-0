package config

import (
	"fmt"
	"strconv"
	"time"
)

const (
	pgHostEnvName         = "PG_HOST"
	pgPortEnvName         = "PG_PORT"
	pgUserEnvName         = "PG_USER"
	pgDatabaseEnvName     = "PG_DB"
	pgMaxOpenConnsEnvName = "PG_MAX_OPEN_CONNS"
	pgMaxIdleConnsEnvName = "PG_MAX_IDLE_CONNS"
	pgMaxIdleTimeEnvName  = "PG_MAX_IDLE_TIME"
)

type PostgresConfig interface {
	DSN() string
	MaxOpenConns() int
	MaxIdleConns() int
	MaxIdleTime() time.Duration
}

type postgresConfig struct {
	host         string
	port         string
	user         string
	db           string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func NewPostgresConfig() (cfg PostgresConfig, err error) {
	host, err := getEnvVariable(pgHostEnvName)
	if err != nil {
		return nil, err
	}

	port, err := getEnvVariable(pgPortEnvName)
	if err != nil {
		return nil, err
	}

	user, err := getEnvVariable(pgUserEnvName)
	if err != nil {
		return nil, err
	}

	db, err := getEnvVariable(pgDatabaseEnvName)
	if err != nil {
		return nil, err
	}

	maxOpenConnsStr, err := getEnvVariable(pgMaxOpenConnsEnvName)
	if err != nil {
		return nil, err
	}
	maxOpenConns, err := strconv.Atoi(maxOpenConnsStr)
	if err != nil {
		return nil, err
	}

	maxIdleConnsStr, err := getEnvVariable(pgMaxIdleConnsEnvName)
	if err != nil {
		return nil, err
	}
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err != nil {
		return nil, err
	}

	maxIdleTimeStr, err := getEnvVariable(pgMaxIdleTimeEnvName)
	if err != nil {
		return nil, err
	}
	maxIdleTime, err := time.ParseDuration(maxIdleTimeStr)
	if err != nil {
		return nil, err
	}

	return &postgresConfig{
		host:         host,
		port:         port,
		user:         user,
		db:           db,
		maxOpenConns: maxOpenConns,
		maxIdleConns: maxIdleConns,
		maxIdleTime:  maxIdleTime,
	}, nil
}

func (cfg *postgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.host,
		cfg.port,
		cfg.user,
		cfg.db,
	)
}

func (cfg *postgresConfig) MaxOpenConns() int {
	return cfg.maxOpenConns
}

func (cfg *postgresConfig) MaxIdleConns() int {
	return cfg.maxIdleConns
}

func (cfg *postgresConfig) MaxIdleTime() time.Duration {
	return cfg.maxIdleTime
}
