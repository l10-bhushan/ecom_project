package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// DbConfig would store the database connection string
type DbConfig struct {
	Dsn string
}

// Config would store port and the database connection string
type Config struct {
	Addr string
	Db   DbConfig
}

// This would be used like a high level wrapper for Config
type Application struct {
	Config Config
}

// Mount method would return us the http handler
// ex : gorilla/mux , chi , gin
// Mount will help us mount all the routes for our api
func (app *Application) Mount() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Server is healthy", http.StatusOK)
	})

	return router
}

// Run function would start our server
func (app *Application) Run(router http.Handler) error {

	// Here we are writing the server configuration
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server up and running on PORT :8080 🚀")

	// Here as we have already configured the server
	// instead of writing http.ListenAndServer(":8080", router)
	// We are writing srv.ListenAndServe()
	// Works the same way
	return srv.ListenAndServe()
}
