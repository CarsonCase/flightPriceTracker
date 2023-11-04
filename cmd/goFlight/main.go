package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/PriceService"
	"github.com/CarsonCase/flightPriceTracker.git/pkg/database"
	"github.com/google/uuid"
)

type priceRoutine struct {
	route   database.Route
	channel chan float64
}

func IterateDates(startDateStr string, endDateStr string, p priceRoutine) {
	// Define the date format
	dateFormat := "01-02-2006"

	// Parse the start and end dates
	startDate, err := time.Parse(dateFormat, startDateStr)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}

	endDate, err := time.Parse(dateFormat, endDateStr)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return
	}

	// Calculate the difference in days
	days := int(endDate.Sub(startDate).Hours() / 24)

	// Iterate through each day
	for i := 0; i <= days; i++ {
		date := startDate.AddDate(0, 0, i)
		go getPrice(p.route.Departure, p.route.Arrival, date.Format(dateFormat), p.channel)
	}
}

func getPrice(departure string, arrival string, date string, c chan float64) {
	price, err := PriceService.GetPrice(departure, arrival, date)
	if err == nil {
		c <- price
	} else {
		c <- 0
	}
}

func goFlight(startDate string, endDate string) {
	routes, err := PriceService.GetRoutes("http://localhost:8000")
	if err != nil {
		fmt.Printf("Error fetching routes: %v", err)
	}

	prices := make(chan float64)

	for _, route := range routes {
		go IterateDates(startDate, endDate, priceRoutine{route, prices})
	}

	for i := range prices {
		fmt.Printf("Price: %v", i)
	}
	close(prices)
}

func postRoute(adminKey string, departure string, arrival string) (int, error) {
	routeParams := database.Route{uuid.UUID{}, departure, arrival}
	requestBody, err := json.Marshal(routeParams)
	if err != nil {
		return 0, err
	}

	// Create a request to the POST /routes endpoint
	req, err := http.NewRequest("POST", "http://localhost:8000/api/routes", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	// Set the 'Authorization' header with the admin key
	req.Header.Set("Authorization", adminKey)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() // Close the response body

	return resp.StatusCode, nil
}

// goFlight commands
// start [start date] [end date]
// new-route [departure] [arrival]
// help
func main() {
	args := os.Args[1:]

	switch args[0] {
	case "help":
		{
			fmt.Println("GoFlight commands\nnew-route [departure] [arrival] publish a new route. Will require an API key which is provided on server startup")
		}
	case "new-route":
		{
			fmt.Println("Enter API Key for goFlight server:")
			reader := bufio.NewReader(os.Stdin)
			apiKey, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			apiKey = strings.TrimRight(apiKey, "\n")

			code, err := postRoute(apiKey, args[1], args[2])

			if err != nil {
				log.Fatal(err)
			}

			if code != http.StatusOK {
				log.Fatal(code)
			}
		}
	default:
		log.Fatal("Invalid selection. Run `goFlight help` for more information")
	}
}
