package main

import (
	"fmt"
	"os"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/PriceService"
)

func getPrice(departure string, arrival string, date string, c chan float64) {
	price, err := PriceService.GetPrice(departure, arrival, date)
	if err != nil {
		fmt.Printf("Error getting price")
	}
	c <- price
}

func goFlight(startDate string, endDate string) {
	routes, err := PriceService.GetRoutes("http://localhost:8000")
	if err != nil {
		fmt.Printf("Error fetching routes: %v", err)
	}

	prices := make(chan float64)

	for _, route := range routes {
		// need a way to for loop through dates :( time.Time probably
		date := "12-12-2023"
		go getPrice(route.Departure, route.Arrival, date, prices)
	}

	x := <-prices
	fmt.Printf("Price: %v", x)
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
			fmt.Println("GoFlight commands\nstart [start date] [end date] - start a new service for all routes between the start and end dates\nnew-route [departure] [arrival] publish a new route. Will require an API key which is provided on server startup")
		}
	case "start":
		{
			goFlight(args[1], args[2])
		}
	case "new-route":
		{

		}
	}
}
