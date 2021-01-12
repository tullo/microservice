package main

import (
	"context"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tullo/microservice/db"
)

func main() {
	var config struct {
		opt     db.ConnectionOptions
		real    bool
		service string
	}
	flag.StringVar(&config.opt.DB.Driver, "db-driver", "mysql", "Database driver")
	flag.StringVar(&config.opt.DB.DSN, "db-dsn", "", "DSN for database connection")
	flag.StringVar(&config.service, "service", "", "Service name for migrations")
	flag.BoolVar(&config.real, "real", false, "false = print migrations, true = run migrations")
	flag.Parse()

	if config.service == "" {
		log.Printf("Available migration services: %+v", db.List())
		log.Fatal()
	}

	ctx := context.Background()

	switch config.real {
	case true:
		handle, err := db.ConnectWithRetry(ctx, config.opt)
		if err != nil {
			log.Fatalf("Error connecting to database: %+v", err)
		}
		if err := db.Run(config.service, handle); err != nil {
			log.Fatalf("An error occured: %+v", err)
		}
	default:
		if err := db.Print(config.service); err != nil {
			log.Fatalf("An error occured: %+v", err)
		}
	}
}
