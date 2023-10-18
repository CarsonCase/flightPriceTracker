package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CarsonCase/flightPriceTracker.git/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

var params database.Flight = database.Flight{
	ID:        uuid.UUID{0},
	Departure: "RNO",
	Arrival:   "SFO",
	Date:      "2023-12-12",
	Price:     104.3,
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

	assert.Equal(t, returnedParams.Price, params.Price)

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
	assert.Equal(t, returnedParams.Price, params.Price)
}

func (c *ApiConfig) getFlight(departure string, arrival string, date string) (*bytes.Buffer, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/flights?departure=JFK&arrival=LAX&date=2023-12-13", nil)
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

func (c *ApiConfig) postFlight(departure string, arrival string, date string) (*bytes.Buffer, int, error) {
	requestBody, err := json.Marshal(params)
	if err != nil {
		return &bytes.Buffer{}, 0, err
	}

	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("POST", "/flights?departure="+departure+"&arrival="+arrival+"&date="+date, bytes.NewBuffer(requestBody))
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
