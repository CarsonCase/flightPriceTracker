package PriceService

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AmadeusAuth struct {
	AccessToken string `json:"access_token"`
}

// Define the Go struct types to represent the JSON data

type Price struct {
	Total string `json:"total"`
}

type Links struct {
	FlightDates  string `json:"flightDates"`
	FlightOffers string `json:"flightOffers"`
}

type FlightDestination struct {
	Type          string `json:"type"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departureDate"`
	ReturnDate    string `json:"returnDate"`
	Price         Price  `json:"price"`
	Links         Links  `json:"links"`
}

type Response struct {
	Data []FlightDestination `json:"data"`
}

func (r *Response) Average() (float64, error) {
	var sum float64
	for _, flight := range r.Data {
		val, err := strconv.ParseFloat(flight.Price.Total, 64)
		if err != nil {
			return 0.0, err
		}
		sum += val
	}
	return (sum / float64(len(r.Data))), nil
}

const oauth2TokenURL = "https://test.api.amadeus.com/v1/security/oauth2/token"

func getPriceEndpoint(departureDest string, arrivalDest string, date string) string {
	return "https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode=" + departureDest + "&destinationLocationCode=" + arrivalDest + "&departureDate=" + date + "&adults=1&nonStop=false&max=250"
}

func sendRequest(req *http.Request) (respBody []byte, err error) {
	// Create an HTTP client
	client := &http.Client{}

	resp, err := client.Do(req)

	if resp.StatusCode != 200 {
		return []byte{}, errors.New("Request Failed. Check inputs.")
	}
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return respBody, err
}

func initAmadeusAuth() (*AmadeusAuth, error) {
	godotenv.Load("../.env")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	apiURL := oauth2TokenURL
	payload := []byte("grant_type=client_credentials&client_id=" + apiKey + "&client_secret=" + apiSecret)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		return &AmadeusAuth{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	respBody, err := sendRequest(req)
	if err != nil {
		return &AmadeusAuth{}, err
	}

	// Parse the response JSON
	var tokenResponse AmadeusAuth
	err = json.Unmarshal(respBody, &tokenResponse)
	if err != nil {
		return &AmadeusAuth{}, err
	}

	return &tokenResponse, err
}
