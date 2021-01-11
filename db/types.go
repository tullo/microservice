package db

import (
	"context"
	"database/sql"
	"time"
)

// DB contains DSN and Driver.
type DB struct {
	DSN    string
	Driver string
}

// ConnectionOptions include common connection options.
type ConnectionOptions struct {
	DB DB

	// Connector is an optional parameter to produce a *sql.DB
	// connection pool, which is then wrapped in *sqlx.DB.
	Connector func(context.Context, DB) (*sql.DB, error)

	Retries        int
	RetryDelay     time.Duration
	ConnectTimeout time.Duration
}
