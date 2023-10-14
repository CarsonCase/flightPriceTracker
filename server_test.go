package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CarsonCase/flightPriceTracker.git/models"

	"github.com/stretchr/testify/assert"
)

var params models.Flight = models.Flight{
	ID:        123,
	Departure: "RNO",
	Arrival:   "SFO",
	Date:      "2023-12-12",
	Price:     104.3,
}

func TestFlight(t *testing.T) {
	// post
	str, code, err := postFlight("RNO", "SFO", "2023-12-12")
	if err != nil {
		t.Fatal("postflight Error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, str.String(), "")

	// get
	str, code, err = getFlight("RNO", "SFO", "2023-12-12")
	if err != nil {
		t.Fatal("getflight Error: ", err.Error())
	}

	decoder := json.NewDecoder(str)

	result := models.Flight{}

	err = decoder.Decode(&result)
	if err != nil {
		t.Fatal("decode error: ", err.Error())
	}

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, code)
	assert.Contains(t, result, params)
}

func getFlight(departure, arrival, date string) (*bytes.Buffer, int, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/flights?departure=JFK&arrival=LAX&date=2023-12-13", nil)
	if err != nil {
		return &bytes.Buffer{}, 0, err
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router := setupRouter() // Replace with your router setup function
	router.ServeHTTP(rr, req)
	return rr.Body, rr.Code, nil
}

func postFlight(departure, arrival, date string) (*bytes.Buffer, int, error) {

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
	router := setupRouter() // Replace with your router setup function
	router.ServeHTTP(rr, req)
	return rr.Body, rr.Code, nil
}
