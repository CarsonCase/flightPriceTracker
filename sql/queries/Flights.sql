-- name: CreateFlight :one
INSERT INTO Flights(ID, created_at, updated_at, departure, arrival, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFlights :one
SELECT * FROM Flights;