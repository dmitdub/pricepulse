package domain

import "time"

type Product struct {
	ID          string
	Name        string
	Description string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
