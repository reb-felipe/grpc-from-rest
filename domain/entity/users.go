package entity

import "time"

type User struct {
	ID          string
	Name        string
	CreatedAt   time.Time
	Coordinates []float64
	UpdatedAt   *time.Time
}
