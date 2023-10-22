package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CarsonCase/flightPriceTracker.git/internal/database"
	"github.com/google/uuid"
)

type Return struct {
	stuff string
}

func (c *ApiConfig) getFlightsHandler(w http.ResponseWriter, r *http.Request) {
	flights, err := c.DB.GetFlights(r.Context())
	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, flights)
}

func (c *ApiConfig) createFlightHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	params := database.Flight{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}

	flight, err := c.DB.CreateFlight(r.Context(), database.CreateFlightParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Route:     params.Route,
		Date:      params.Date,
		Price:     params.Price,
	})

	if err != nil {
		RespondWithError(w, 400, "couldn't create flight"+err.Error())
	}

	RespondWithJson(w, http.StatusOK, flight)
}
