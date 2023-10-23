package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

var routeParams database.Route = database.Route{
	ID:        uuid.UUID{0},
	Departure: "RNO",
	Arrival:   "SFO",
}

var flightParams database.Flight = database.Flight{
	ID:    uuid.UUID{0},
	Route: uuid.UUID{0},
	Date:  "2023-12-12",
	Price: 104.3,
}

// Happy path does this
// 1. Post a new route
// 2. Get that route ID
// 3. Post a flgiht to that ID
// 4. Get that flight
func TestFlightHappyPath(t *testing.T) {
	// Set up Database connection
	godotenv.Load()
	sqlString := os.Getenv("DB_URL")
	connection, err := sql.Open("postgres", sqlString)

	if err != nil {
		t.Fatal("SQL error:", err)
	}

	apiCfg := ApiConfig{
		DB: database.New(connection),
	}

	router := apiCfg.setupRouter()

	// post a new route
	code, err := postRoute(router, apiCfg.adminKey, "RNO", "SFO")
	if err != nil {
		t.Fatal("postRoute Error: ", err.Error())
	}

	// Check the status code and response body. pass if 200
	assert.Equal(t, http.StatusOK, code)

	// get list of routes
	str, code, err := getRoutes(router)
	if err != nil {
		t.Fatal("postRoute Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

	// convert string response to []Route
	returnedRoutesParams := []database.Route{}
	err = json.Unmarshal([]byte(str), &returnedRoutesParams)
	if err != nil {
		t.Fatal("Unmarshal error: ", err)
	}

	// get latest routeID
	latestRoute := returnedRoutesParams[0]

	// post a flight
	str, code, err = postFlight(router, apiCfg.adminKey, latestRoute.ID, "2023-12-12")
	if err != nil {
		t.Fatal("postflight Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

	returnedFlightParams := database.Flight{}
	err = json.Unmarshal([]byte(str), &returnedFlightParams)
	if err != nil {
		t.Fatal("Unmarshal error: ", err)
	}

	assert.Equal(t, returnedFlightParams.Price, flightParams.Price)

	// get the list of flights for that route
	str, code, err = getFlights(router, latestRoute.ID)
	if err != nil {
		t.Fatal("getflight Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)

	returnedFlightsParams := []database.Flight{}
	err = json.Unmarshal([]byte(str), &returnedFlightsParams)
	if err != nil {
		t.Fatal("Unmarshal error: ", err)
	}

	assert.Equal(t, returnedFlightsParams[0].Price, flightParams.Price)
}

///==================
/// HELPER FUNCTIONS
///==================

// Helper function to get Routes
func getRoutes(router *chi.Mux) (string, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/routes", nil)
	if err != nil {
		return "", 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router.ServeHTTP(rr, req)
	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return "", 0, err
	}
	return string(str), rr.Code, nil
}

// helper function to post a route
func postRoute(router *chi.Mux, adminKey string, departure string, arrival string) (int, error) {
	// if route already exists don't post it again
	if routeParams.Arrival == arrival && routeParams.Departure == departure {
		return http.StatusOK, nil
	}

	routeParams.Departure = departure
	routeParams.Arrival = arrival
	requestBody, err := json.Marshal(routeParams)
	if err != nil {
		return 0, err
	}

	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("POST", "/api/routes", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	req.Header.Add("Authorization", adminKey)

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create API router and handle the request
	router.ServeHTTP(rr, req)
	return rr.Code, nil
}

// helper function to get flights
func getFlights(router *chi.Mux, routeID uuid.UUID) (string, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/flights", nil)
	if err != nil {
		return "", 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router.ServeHTTP(rr, req)
	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return "", 0, err
	}
	return string(str), rr.Code, nil
}

// helper functo to post a flight
func postFlight(router *chi.Mux, adminKey string, routeID uuid.UUID, date string) (string, int, error) {
	flightParams.Route = routeID
	requestBody, err := json.Marshal(flightParams)
	if err != nil {
		return "", 0, err
	}

	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("POST", "/api/flights", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", 0, err
	}

	req.Header.Add("Authorization", adminKey)

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router.ServeHTTP(rr, req)
	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return "", 0, err
	}
	return string(str), rr.Code, nil
}
