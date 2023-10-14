package main

import "net/http"

type Return struct {
	stuff string
}

func flightsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("expected response content"))
	// RespondWithJson(w, 200, Return{"expected response content"})
}
