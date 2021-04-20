package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// Connect connects to a database and produces the handle for injection.
func Connect(ctx context.Context) (*sqlx.DB, error) {
	dsn := os.Getenv("DB_DSN")
	driver := os.Getenv("DB_DRIVER")
	if dsn == "" {
		return nil, fmt.Errorf("db_dsn not provided")
	}
	if driver == "" {
		driver = "mysql"
	}

	opt := ConnectionOptions{
		Retries:        5,
		RetryDelay:     5 * time.Second,
		ConnectTimeout: 5 * time.Minute,
		DB: DB{
			DSN:    dsn,
			Driver: driver,
		},
		Connector: Connector,
	}

	return ConnectWithOptions(ctx, opt)
}

// ConnectWithOptions connects to a database host using connection options.
func ConnectWithOptions(ctx context.Context, options ConnectionOptions) (*sqlx.DB, error) {
	db := options.DB
	if db.DSN == "" {
		return nil, fmt.Errorf("dsn not provided")
	}
	if db.Driver == "" {
		db.Driver = "mysql"
	}
	db.DSN = cleanDSN(db.DSN)

	con, err := connect(ctx, options)
	if err != nil {
		return nil, err
	}
	con.SetMaxOpenConns(800)
	con.SetMaxIdleConns(800)

	return con, nil
}

func connect(ctx context.Context, options ConnectionOptions) (*sqlx.DB, error) {
	db := options.DB
	if options.Connector != nil {
		handle, err := options.Connector(ctx, db)
		if err != nil {
			return nil, fmt.Errorf("error producing connection pool: %w", err)
		}

		return sqlx.NewDb(handle, db.Driver), nil
	}

	ping, err := sqlx.ConnectContext(ctx, db.Driver, db.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return ping, nil
}
