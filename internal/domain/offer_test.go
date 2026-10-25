package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewOffer(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name          string
		id            string
		productID     string
		regularPrice  int64
		currentPrice  int64
		currency      string
		stockQuantity int
		createdAt     time.Time
		wantErr       error
	}{
		{
			name:          "valid offer",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       nil,
		},
		{
			name:          "empty offer ID",
			id:            "",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidOfferID,
		},
		{
			name:          "empty product ID",
			id:            "offer-1",
			productID:     "",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidProductID,
		},
		{
			name:          "negative regular price",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  -100,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidRegularPrice,
		},
		{
			name:          "zero regular price",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  0,
			currentPrice:  0,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidRegularPrice,
		},
		{
			name:          "negative current price",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  -100,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidCurrentPrice,
		},
		{
			name:          "invalid currency length",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EU",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidCurrency,
		},
		{
			name:          "lowercase currency",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "eur",
			stockQuantity: 10,
			createdAt:     createdAt,
			wantErr:       ErrInvalidCurrency,
		},
		{
			name:          "negative stock quantity",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: -1,
			createdAt:     createdAt,
			wantErr:       ErrInvalidStockQuantity,
		},
		{
			name:          "zero creation time",
			id:            "offer-1",
			productID:     "product-1",
			regularPrice:  10000,
			currentPrice:  8000,
			currency:      "EUR",
			stockQuantity: 10,
			createdAt:     time.Time{},
			wantErr:       ErrInvalidCreatedAt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer, err := NewOffer(
				tt.id,
				tt.productID,
				tt.regularPrice,
				tt.currentPrice,
				tt.currency,
				tt.stockQuantity,
				tt.createdAt,
			)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("NewOffer() error = %v, want %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			if offer.ID() != tt.id {
				t.Errorf("ID() = %q, want %q", offer.ID(), tt.id)
			}

			if offer.CurrentPrice() != tt.currentPrice {
				t.Errorf("CurrentPrice() = %d, want %d", offer.CurrentPrice(), tt.currentPrice)
			}

			if offer.StockQuantity() != tt.stockQuantity {
				t.Errorf("StockQuantity() = %d, want %d", offer.StockQuantity(), tt.stockQuantity)
			}
		})
	}
}

func TestOfferDiscountPercent(t *testing.T) {
	tests := []struct {
		name         string
		regularPrice int64
		currentPrice int64
		want         int
	}{
		{
			name:         "20 percent discount",
			regularPrice: 10000,
			currentPrice: 8000,
			want:         20,
		},
		{
			name:         "50 percent discount",
			regularPrice: 10000,
			currentPrice: 5000,
			want:         50,
		},
		{
			name:         "no discount",
			regularPrice: 10000,
			currentPrice: 10000,
			want:         0,
		},
		{
			name:         "current price above regular price",
			regularPrice: 10000,
			currentPrice: 12000,
			want:         0,
		},
		{
			name:         "free product",
			regularPrice: 10000,
			currentPrice: 0,
			want:         100,
		},
		{
			name:         "zero value offer",
			regularPrice: 0,
			currentPrice: 0,
			want:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer := Offer{
				id:            "",
				productID:     "",
				regularPrice:  tt.regularPrice,
				currentPrice:  tt.currentPrice,
				currency:      "",
				stockQuantity: 0,
				createdAt:     time.Time{},
				updatedAt:     time.Time{},
			}

			got := offer.DiscountPercent()

			if got != tt.want {
				t.Errorf("DiscountPercent() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestOfferUpdatePrice(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)

	tests := []struct {
		name        string
		newPrice    int64
		updatedAt   time.Time
		wantErr     error
		wantPrice   int64
		wantUpdated time.Time
	}{
		{
			name:        "valid price update",
			newPrice:    7000,
			updatedAt:   updatedAt,
			wantErr:     nil,
			wantPrice:   7000,
			wantUpdated: updatedAt,
		},
		{
			name:        "zero price",
			newPrice:    0,
			updatedAt:   updatedAt,
			wantErr:     nil,
			wantPrice:   0,
			wantUpdated: updatedAt,
		},
		{
			name:        "negative price",
			newPrice:    -100,
			updatedAt:   updatedAt,
			wantErr:     ErrInvalidCurrentPrice,
			wantPrice:   8000,
			wantUpdated: createdAt,
		},
		{
			name:        "zero update time",
			newPrice:    7000,
			updatedAt:   time.Time{},
			wantErr:     ErrInvalidUpdatedAt,
			wantPrice:   8000,
			wantUpdated: createdAt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer, err := NewOffer(
				"offer-1",
				"product-1",
				10000,
				8000,
				"EUR",
				10,
				createdAt,
			)
			if err != nil {
				t.Fatalf("NewOffer() error = %v", err)
			}

			err = offer.UpdatePrice(tt.newPrice, tt.updatedAt)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("UpdatePrice() error = %v, want %v", err, tt.wantErr)
			}

			if offer.CurrentPrice() != tt.wantPrice {
				t.Errorf("CurrentPrice() = %d, want %d", offer.CurrentPrice(), tt.wantPrice)
			}

			if !offer.UpdatedAt().Equal(tt.wantUpdated) {
				t.Errorf("UpdatedAt() = %v, want %v", offer.UpdatedAt(), tt.wantUpdated)
			}
		})
	}
}

func TestOfferUpdateStock(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)

	tests := []struct {
		name         string
		newQuantity  int
		updatedAt    time.Time
		wantErr      error
		wantQuantity int
		wantUpdated  time.Time
	}{
		{
			name:         "valid stock update",
			newQuantity:  5,
			updatedAt:    updatedAt,
			wantErr:      nil,
			wantQuantity: 5,
			wantUpdated:  updatedAt,
		},
		{
			name:         "zero stock",
			newQuantity:  0,
			updatedAt:    updatedAt,
			wantErr:      nil,
			wantQuantity: 0,
			wantUpdated:  updatedAt,
		},
		{
			name:         "negative stock",
			newQuantity:  -1,
			updatedAt:    updatedAt,
			wantErr:      ErrInvalidStockQuantity,
			wantQuantity: 10,
			wantUpdated:  createdAt,
		},
		{
			name:         "zero update time",
			newQuantity:  5,
			updatedAt:    time.Time{},
			wantErr:      ErrInvalidUpdatedAt,
			wantQuantity: 10,
			wantUpdated:  createdAt,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer, err := NewOffer(
				"offer-1",
				"product-1",
				10000,
				8000,
				"EUR",
				10,
				createdAt,
			)
			if err != nil {
				t.Fatalf("NewOffer() error = %v", err)
			}

			err = offer.UpdateStock(tt.newQuantity, tt.updatedAt)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("UpdateStock() error = %v, want %v", err, tt.wantErr)
			}

			if offer.StockQuantity() != tt.wantQuantity {
				t.Errorf("StockQuantity() = %d, want %d", offer.StockQuantity(), tt.wantQuantity)
			}

			if !offer.UpdatedAt().Equal(tt.wantUpdated) {
				t.Errorf("UpdatedAt() = %v, want %v", offer.UpdatedAt(), tt.wantUpdated)
			}
		})
	}
}
