-- name: CreateRoute :one
INSERT INTO Routes(id, departure, arrival)
VALUES($1, $2, $3)
RETURNING *;

-- name: GetRoutes :many
SELECT * FROM Routes;
