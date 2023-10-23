package main

import (
	"encoding/json"
	"net/http"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/database"
	"github.com/google/uuid"
)

func (c *ApiConfig) getRoutesHandler(w http.ResponseWriter, r *http.Request) {
	routes, err := c.DB.GetRoutes(r.Context())
	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}
	RespondWithJson(w, 200, routes)
}

func (c *ApiConfig) createRouteHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	params := database.Route{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, err.Error())
		return
	}

	flight, err := c.DB.CreateRoute(r.Context(), database.CreateRouteParams{
		ID:        uuid.New(),
		Departure: params.Departure,
		Arrival:   params.Arrival,
	})

	if err != nil {
		RespondWithError(w, 400, "couldn't create route"+err.Error())
	}

	RespondWithJson(w, http.StatusOK, flight)
}
