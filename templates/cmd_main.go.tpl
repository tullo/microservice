package main

// file is autogenerated, do not modify here, see
// generator and template: templates/cmd_main.go.tpl

import (
	"log"

	"net/http"

	"github.com/SentimensRG/sigctx"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"${MODULE}/db"
	"${MODULE}/internal"
	"${MODULE}/rpc/${SERVICE}"
	server "${MODULE}/server/${SERVICE}"
)

func main() {
	var config struct {
		migrate bool
		opt     db.ConnectionOptions
	}
	flag.StringVar(&config.opt.DB.Driver, "migrate-db-driver", "mysql", "Migrations: Database driver")
	flag.StringVar(&config.opt.DB.DSN, "migrate-db-dsn", "", "Migrations: DSN for database connection")
	flag.BoolVar(&config.migrate, "migrate", false, "Run migrations?")
	flag.Parse()

	ctx := sigctx.New()

	if config.migrate {
		handle, err := db.ConnectWithRetry(ctx, config.opt)
		if err != nil {
			log.Fatalf("Error connecting to database: %+v", err)
		}
		if err := db.Run("${SERVICE}", handle); err != nil {
			log.Fatalf("An error occurred: %+v", err)
		}
	}

	srv, err := server.New(ctx)
	if err != nil {
		log.Fatalf("Error in service.New(): %+v", err)
	}

	twirpServer := ${SERVICE}.New${SERVICE_CAMEL}ServiceServer(srv, internal.NewServerHooks())

	go func() {
		log.Println("Starting service on port :3000")
		err := http.ListenAndServe(":3000", internal.WrapAll(twirpServer))
		if !errors.Is(err, http.ErrServerClosed) {
			log.Println("Server error:", err)
		}
	}()
	<-ctx.Done()

	srv.Shutdown()
	log.Println("Done.")
}