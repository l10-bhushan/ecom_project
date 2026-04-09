package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/l10-bhushan/ecom_project/internal/env"
	"github.com/l10-bhushan/ecom_project/internal/router"
)

func main() {
	ctx := context.Background()
	// Here we are creating instance of Config struct adding the Address
	cfg := router.Config{
		Addr: ":8080",
		Db: router.DbConfig{
			Dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom_db sslmode=disable"),
		},
	}

	// Setting up the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	// Connecting to db
	conn, err := pgx.Connect(ctx, cfg.Db.Dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.Db.Dsn)

	// Then we are creating the Application struct that will initialize our
	// router using the config object
	app := router.Application{
		Config: cfg,
	}

	// Initializaing the service

	// Initializing the handler
	// handler := products.Handler{}

	// Mount method on router will mount all the routes into the router
	// Run method would start the server
	if err := app.Run(app.Mount()); err != nil {
		slog.Error("Server failed to start", "error", nil)
		log.Printf("Server has failed to start: %s", err)
		os.Exit(0)
	}
}
