package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func initAmadeusAuth(apiKey string, apiSecret string) *AmadeusAuth {
	apiURL := "https://test.api.amadeus.com/v1/security/oauth2/token"
	payload := []byte("grant_type=client_credentials&client_id=" + apiKey + "&client_secret=" + apiSecret)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return &AmadeusAuth{}
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return &AmadeusAuth{}
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return &AmadeusAuth{}
	}

	// Parse the response JSON
	var tokenResponse AmadeusAuth
	err = json.Unmarshal(respBody, &tokenResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return &AmadeusAuth{}
	}

	return &tokenResponse
}

func getPrice(departureDest, arrivalDest string) (float64, error) {
	auth := initAmadeusAuth("apiKey", "apiSecret")

	req, err := http.NewRequest("GET", "https://test.api.amadeus.com/v2/shopping/flight-offers?originLocationCode="+departureDest+"&destinationLocationCode="+arrivalDest+"&departureDate=2023-12-01&adults=1&nonStop=false&max=250", nil)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0.0, err
	}

	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)

	// Create an HTTP client
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0.0, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return 0.0, err
	}

	// Unmarshal JSON data into the defined struct
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0.0, err
	}

	val, err := strconv.ParseFloat(response.Data[0].Price.Total, 64)
	if err != nil {
		fmt.Println("Error converting output price string:", err)
		return 0.0, err
	}
	return val, nil
}

func main() {

}
