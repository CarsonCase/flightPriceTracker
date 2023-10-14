// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Flight struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Departure string
	Arrival   string
	Price     float64
}
