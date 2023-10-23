package main

import (
	"fmt"

	"github.com/CarsonCase/flightPriceTracker.git/pkg/PriceService"
)

// goFlight commands
// start [start date] [end date]
// new-route [departure] [arrival]
// help
func main() {
	// args := os.Args[1:]

	// get list of routes
	routes, err := PriceService.GetRoutes("http://localhost:8000")
	if err != nil {
		fmt.Printf("Error fetching routes: %v", err)
	}

	fmt.Println(routes)

	// got, _ := getPrice(args[0], args[1], args[2])

	// fmt.Printf("RNO -> SF: %v\n", (got))
}
