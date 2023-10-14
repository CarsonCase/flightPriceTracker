package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFlights(t *testing.T) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", "/flights?departure=JFK&arrival=LAX&date=2023-12-13", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create your API router and handle the request
	router := setupRouter() // Replace with your router setup function
	router.ServeHTTP(rr, req)

	// Check the status code and response body
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "expected response content")
}
