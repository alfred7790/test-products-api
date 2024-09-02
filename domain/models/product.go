package models

import "time"

type Product struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Code        string
	Name        string
	Price       float64
	Description string
}
