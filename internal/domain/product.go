package domain

import (
	"strings"
	"time"
)

type Product struct {
	id          string
	name        string
	description string
	category    string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewProduct(
	id string,
	name string,
	description string,
	category string,
	createdAt time.Time,
) (Product, error) {
	product := Product{
		id:          id,
		name:        name,
		description: description,
		category:    category,
		createdAt:   createdAt,
		updatedAt:   createdAt,
	}

	if err := product.Validate(); err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p Product) Validate() error {
	if strings.TrimSpace(p.id) == "" {
		return ErrInvalidProductID
	}

	if strings.TrimSpace(p.name) == "" {
		return ErrInvalidProductName
	}

	if strings.TrimSpace(p.description) == "" {
		return ErrInvalidProductDescription
	}

	if strings.TrimSpace(p.category) == "" {
		return ErrInvalidProductCategory
	}

	if p.createdAt.IsZero() {
		return ErrInvalidProductCreatedAt
	}

	if p.updatedAt.IsZero() {
		return ErrInvalidProductUpdatedAt
	}

	return nil
}

func (p Product) ID() string {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Description() string {
	return p.description
}

func (p Product) Category() string {
	return p.category
}

func (p Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p Product) UpdatedAt() time.Time {
	return p.updatedAt
}
