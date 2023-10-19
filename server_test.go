package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CarsonCase/flightPriceTracker.git/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

var routeParams database.Route = database.Route{
	ID:        uuid.UUID{1},
	Departure: "RNO",
	Arrival:   "SFO",
}

var flightParams database.Flight = database.Flight{
	ID:    uuid.UUID{0},
	Route: uuid.UUID{1},
	Date:  "2023-12-12",
	Price: 104.3,
}

// Todo
// Find out why post isn't working the second time...
// Combine the post and get route functions in happy path
// Make happy path do this
// 1. Post a new route
// 2. Get that route ID
// 3. Post a flgiht to that ID
// 4. Get that flight

func _TestPostRoute(t *testing.T) {
	godotenv.Load()
	sqlString := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", sqlString)

	if err != nil {
		t.Fatal("SQL error:", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(connection),
	}

	// post a route
	code, err := apiCfg.postRoute("RNO", "SFO")
	if err != nil {
		t.Fatal("postRoute Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)
}

func TestGetRoute(t *testing.T) {
	godotenv.Load()
	sqlString := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", sqlString)
	if err != nil {
		t.Fatal("SQL error:", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(connection),
	}

	// get routes
	str, code, err := apiCfg.getRoutes()
	if err != nil {
		t.Fatal("postRoute Error: ", err.Error())
	}

	fmt.Println(str)

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

}

func (c *ApiConfig) getRoutes() (string, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/routes", nil)
	if err != nil {
		return "", 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router := c.setupRouter()
	router.ServeHTTP(rr, req)
	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return "", 0, err
	}
	return string(str), rr.Code, nil
}

func (c *ApiConfig) postRoute(departure string, arrival string) (int, error) {
	requestBody, err := json.Marshal(routeParams)
	if err != nil {
		return 0, err
	}

	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("POST", "/routes", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create API router and handle the request
	router := c.setupRouter() // Replace with your router setup function
	router.ServeHTTP(rr, req)
	return rr.Code, nil
}

func FlightHappyPath(t *testing.T) {
	// Set up
	godotenv.Load()
	sqlString := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", sqlString)

	if err != nil {
		t.Fatal("SQL error:", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(connection),
	}

	// post a flight
	str, code, err := apiCfg.postFlight("RNO", "SFO", "2023-12-12")
	if err != nil {
		t.Fatal("postflight Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

	returnedParams := database.Flight{}
	err = json.Unmarshal(str.Bytes(), &returnedParams)
	if err != nil {
		t.Fatal("Unmarshal error: ", err)
	}

	assert.Equal(t, returnedParams.Price, flightParams.Price)

	// get the flight back
	str, code, err = apiCfg.getFlight("RNO", "SFO", "2023-12-12")
	if err != nil {
		t.Fatal("getflight Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

	err = json.Unmarshal(str.Bytes(), &returnedParams)
	if err != nil {
		t.Fatal("Unmarshal error: ", err)
	}

	assert.Equal(t, returnedParams.Price, flightParams.Price)
}

func (c *ApiConfig) getFlight(departure string, arrival string, date string) (*bytes.Buffer, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/flights", nil)
	if err != nil {
		return &bytes.Buffer{}, 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router := c.setupRouter()
	router.ServeHTTP(rr, req)
	return rr.Body, rr.Code, nil
}

func (c *ApiConfig) postFlight(departure string, arrival string, date string) (*bytes.Buffer, int, error) {
	requestBody, err := json.Marshal(flightParams)
	if err != nil {
		return &bytes.Buffer{}, 0, err
	}

	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("POST", "/flights", bytes.NewBuffer(requestBody))
	if err != nil {
		return &bytes.Buffer{}, 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router := c.setupRouter() // Replace with your router setup function
	router.ServeHTTP(rr, req)
	return rr.Body, rr.Code, nil
}
