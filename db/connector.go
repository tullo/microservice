package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// apm specific wrapper for the go mysql driver
	_ "go.elastic.co/apm/module/apmsql/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.elastic.co/apm/module/apmsql"
)

// Connector opens a database and pings it using APM SQL wrapper.
func Connector(ctx context.Context, opt DB) (*sql.DB, error) {
	log.Println("Connector: apmsql.Open(driver, dsn)", opt.Driver, maskDSN(opt.DSN))
	db, err := apmsql.Open(opt.Driver, opt.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database: %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		db.Close()

		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}

// ConnectWithRetry uses retry options set in ConnectionOptions.
func ConnectWithRetry(ctx context.Context, options ConnectionOptions) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)
	// mask db credentials ready for output in logs and such.
	dsn := maskDSN(options.DB.DSN)

	// by default, retry for 5 minutes, 5 seconds between retries.
	if options.Retries == 0 && options.ConnectTimeout.Seconds() == 0 {
		options.ConnectTimeout = 5 * time.Minute
		options.RetryDelay = 5 * time.Second
	}

	connErrCh := make(chan error, 1)
	defer close(connErrCh)

	log.Println("connecting to database", dsn)

	go func() {
		try := 0
		for {
			try++
			if options.Retries > 0 && options.Retries <= try {
				err = errors.Errorf("could not connect, dsn=%s, tries=%d", dsn, try)

				break
			}

			db, err = ConnectWithOptions(ctx, options)
			if err != nil {
				log.Printf("can't connect, dsn=%s, err=%s, try=%d", dsn, err, try)

				select {
				case <-ctx.Done():
					break
				case <-time.After(options.RetryDelay):
					continue
				}
			}

			break
		}
		// run out of retries; signal error.
		connErrCh <- err
	}()

	select {
	case err = <-connErrCh:
		break
	case <-time.After(options.ConnectTimeout):
		return nil, errors.Errorf("db connect timed out, dsn=%s", dsn)
	case <-ctx.Done():
		return nil, errors.Errorf("db connection cancelled, dsn=%s", dsn)
	}

	return db, nil
}
