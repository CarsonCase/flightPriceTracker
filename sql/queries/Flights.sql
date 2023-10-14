-- name: CreateFlight :one
INSERT INTO Flights(ID, created_at, updated_at, departure, arrival, date, price)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFlights :one
SELECT * FROM Flights;