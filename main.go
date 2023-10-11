package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getPrice(departureDest string, arrivalDest string, date string) (float64, error) {
	// make request to getPrice endpoint with auth token
	auth, err := initAmadeusAuth()

	if err != nil {
		fmt.Println("Error generating auth token:", err)
		return 0.0, err
	}
	req, err := http.NewRequest("GET", getPriceEndpoint(departureDest, arrivalDest, date), nil)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0.0, err
	}

	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)

	respBody, err := sendRequest(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0.0, err
	}

	// Unmarshal JSON data into the defined struct
	var response Response

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return 0.0, err
	}

	// get lowest price and format
	val, err := response.Average()
	if err != nil {
		fmt.Println("Error converting output price string:", err)
		return 0.0, err
	}
	// return
	return val, nil
}

func main() {
	args := os.Args[1:]
	got, _ := getPrice(args[0], args[1], args[2])

	fmt.Printf("RNO -> SF: %v\n", (got))
}
