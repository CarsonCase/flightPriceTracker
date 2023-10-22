package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
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
	DB       *database.Queries
	adminKey string
}

func (c *ApiConfig) setupRouter() *chi.Mux {
	// generate a new admin key
	c.adminKey = generateAdminKey()

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

	adminRouter := chi.NewRouter()
	adminRouter.Use(c.AuthMiddleware)

	adminRouter.Post("/flights", c.createFlightHandler)
	adminRouter.Post("/routes", c.createRouteHandler)
	router.Mount("/api", adminRouter)

	return router
}

func generateAdminKey() string {
	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		// Handle the error
	}
	apiKey := hex.EncodeToString(keyBytes)
	return apiKey
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
	log.Printf("Admin API key: %v\n", apiCfg.adminKey)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
