package PriceService

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/CarsonCase/flightPriceTracker.git/internal/database"
)

func GetRoutes(url string) ([]database.Route, error) {
	// Create a request to the GET /flights endpoint
	req, err := http.NewRequest("GET", url+"/routes", nil)
	if err != nil {
		return []database.Route{}, nil
	}

	client := &http.Client{}

	rr, err := client.Do(req)
	if err != nil {
		return []database.Route{}, nil
	}

	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return []database.Route{}, nil
	}

	// convert string response to []Route
	returnedRoutesParams := []database.Route{}
	err = json.Unmarshal([]byte(str), &returnedRoutesParams)
	if err != nil {
		return []database.Route{}, nil
	}

	return returnedRoutesParams, nil
}
