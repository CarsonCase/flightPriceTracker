-- name: CreateFlight :one
INSERT INTO Flights(ID, created_at, updated_at, route, date, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFlights :many
SELECT * FROM Flights ORDER BY created_at DESC;