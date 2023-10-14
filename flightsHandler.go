package main

import "net/http"

type Return struct {
	stuff string
}

func getFlightsHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, 200, Return{""})
}

func createFlightHandler(w http.ResponseWriter, r *http.Request) {

}
