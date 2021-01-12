package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	// apm specific wrapper for the go mysql driver
	_ "go.elastic.co/apm/module/apmsql/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.elastic.co/apm/module/apmsql"
)

// Connect connects to a database and produces the handle for injection.
func Connect(ctx context.Context) (*sqlx.DB, error) {

	dsn := os.Getenv("DB_DSN")
	driver := os.Getenv("DB_DRIVER")
	if dsn == "" {
		return nil, errors.New("DB_DSN not provided")
	}
	if driver == "" {
		driver = "mysql"
	}

	options := ConnectionOptions{
		Retries:        5,
		RetryDelay:     5 * time.Second,
		ConnectTimeout: 5 * time.Minute,
		DB: DB{
			DSN:    dsn,
			Driver: driver,
		},
		Connector: func(ctx context.Context, opt DB) (*sql.DB, error) {
			db, err := apmsql.Open(opt.Driver, opt.DSN)
			if err != nil {
				return nil, fmt.Errorf("failed to open a database: %w", err)
			}
			if err = db.PingContext(ctx); err != nil {
				db.Close()

				return nil, fmt.Errorf("failed to ping the database: %w", err)
			}

			return db, nil
		},
	}

	return ConnectWithOptions(ctx, options)
}

// ConnectWithOptions connects to a database host using connection options.
func ConnectWithOptions(ctx context.Context, options ConnectionOptions) (*sqlx.DB, error) {
	db := options.DB
	if db.DSN == "" {
		return nil, errors.New("DSN not provided")
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
		if err == nil {
			return sqlx.NewDb(handle, db.Driver), nil
		}
		return nil, errors.WithStack(err)
	}

	return sqlx.ConnectContext(ctx, db.Driver, db.DSN)
}
