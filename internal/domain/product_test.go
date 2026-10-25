package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewProduct(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name        string
		id          string
		nameValue   string
		description string
		category    string
		createdAt   time.Time
		wantErr     error
	}{
		{
			name:        "valid product",
			id:          "product-1",
			nameValue:   "Smartphone",
			description: "Modern smartphone",
			category:    "Electronics",
			createdAt:   createdAt,
			wantErr:     nil,
		},
		{
			name:        "empty product ID",
			id:          "",
			nameValue:   "Smartphone",
			description: "Modern smartphone",
			category:    "Electronics",
			createdAt:   createdAt,
			wantErr:     ErrInvalidProductID,
		},
		{
			name:        "empty product name",
			id:          "product-1",
			nameValue:   "",
			description: "Modern smartphone",
			category:    "Electronics",
			createdAt:   createdAt,
			wantErr:     ErrInvalidProductName,
		},
		{
			name:        "empty product description",
			id:          "product-1",
			nameValue:   "Smartphone",
			description: "",
			category:    "Electronics",
			createdAt:   createdAt,
			wantErr:     ErrInvalidProductDescription,
		},
		{
			name:        "empty product category",
			id:          "product-1",
			nameValue:   "Smartphone",
			description: "Modern smartphone",
			category:    "",
			createdAt:   createdAt,
			wantErr:     ErrInvalidProductCategory,
		},
		{
			name:        "zero creation time",
			id:          "product-1",
			nameValue:   "Smartphone",
			description: "Modern smartphone",
			category:    "Electronics",
			createdAt:   time.Time{},
			wantErr:     ErrInvalidProductCreatedAt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(
				tt.id,
				tt.nameValue,
				tt.description,
				tt.category,
				tt.createdAt,
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewProduct() error = %v, want %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			if product.ID() != tt.id {
				t.Errorf("ID() = %q, want %q", product.ID(), tt.id)
			}

			if product.Name() != tt.nameValue {
				t.Errorf("Name() = %q, want %q", product.Name(), tt.nameValue)
			}

			if product.Description() != tt.description {
				t.Errorf("Description() = %q, want %q", product.Description(), tt.description)
			}

			if product.Category() != tt.category {
				t.Errorf("Category() = %q, want %q", product.Category(), tt.category)
			}

			if !product.CreatedAt().Equal(createdAt) {
				t.Errorf("CreatedAt() = %v, want %v", product.CreatedAt(), createdAt)
			}

			if !product.UpdatedAt().Equal(createdAt) {
				t.Errorf("UpdatedAt() = %v, want %v", product.UpdatedAt(), createdAt)
			}
		})
	}
}
