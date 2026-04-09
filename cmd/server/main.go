package main

import (
	"log"
	"os"

	"github.com/l10-bhushan/ecom_project/internal/router"
)

func main() {
	// Here we are creating instance of Config struct adding the Address
	cfg := router.Config{
		Addr: ":8080",
	}

	// Then we are creating the Application struct that will initialize our
	// router using the config object
	app := router.Application{
		Config: cfg,
	}

	// Mount method on router will mount all the routes into the router

	router := app.Mount()

	// Run method would start the server
	if err := app.Run(router); err != nil {
		log.Printf("Server has failed to start: %s", err)
		os.Exit(0)
	}
}
