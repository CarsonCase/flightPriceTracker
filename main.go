package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CarsonCase/flightPriceTracker.git/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func (c *ApiConfig) setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Get("/flights", c.getFlightsHandler)
	router.Get("/routes", c.getRoutesHandler)
	router.Post("/flights", c.createFlightHandler)
	router.Post("/routes", c.createRouteHandler)
	return router
}

func main() {
	godotenv.Load()
	portString := os.Getenv("SERVER_PORT")

	sqlString := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", sqlString)

	if err != nil {
		log.Fatal("SQL error: ", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(connection),
	}

	router := apiCfg.setupRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v\n", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(portString)
}
