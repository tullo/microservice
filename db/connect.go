package db

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	var options ConnectionOptions
	options.DB.DSN = dsn
	options.DB.Driver = driver

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

	if options.Connector != nil {
		handle, err := options.Connector(ctx, db)
		if err == nil {
			return sqlx.NewDb(handle, db.Driver), nil
		}
		return nil, errors.WithStack(err)
	}

	return sqlx.ConnectContext(ctx, db.Driver, db.DSN)
}
