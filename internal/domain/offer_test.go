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
			name: "zero value offer",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offer := Offer{
				regularPrice: tt.regularPrice,
				currentPrice: tt.currentPrice,
			}

			got := offer.DiscountPercent()

			if got != tt.want {
				t.Errorf("DiscountPercent() = %d, want %d", got, tt.want)
			}
		})
	}
}
