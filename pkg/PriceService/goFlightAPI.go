package PriceService

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/database"
)

func GetRoutes(url string) ([]database.Route, error) {
	req, err := http.NewRequest("GET", url+"/routes", nil)
	if err != nil {
		return []database.Route{}, err
	}

	client := &http.Client{}

	rr, err := client.Do(req)
	if err != nil {
		return []database.Route{}, err
	}

	defer rr.Body.Close() // Close the response body to prevent leaks

	str, err := io.ReadAll(rr.Body)
	if err != nil {
		return []database.Route{}, err
	}

	// convert string response to []Route
	returnedRoutesParams := []database.Route{}
	err = json.Unmarshal([]byte(str), &returnedRoutesParams)
	if err != nil {
		return []database.Route{}, err
	}

	return returnedRoutesParams, nil
}
