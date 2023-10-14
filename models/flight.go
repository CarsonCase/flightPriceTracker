package models

// Flight represents the flight data model.
type Flight struct {
	ID        int     `json:"id"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
	Date      string  `json:"date"`
	Price     float64 `json:"price"`
}
