package domain

import (
	"strings"
	"time"
)

const CurrencyCodeLength = 3

type Offer struct {
	id            string
	productID     string
	regularPrice  int64
	currentPrice  int64
	currency      string
	stockQuantity int
	createdAt     time.Time
	updatedAt     time.Time
}

func NewOffer(
	id string,
	productID string,
	regularPrice int64,
	currentPrice int64,
	currency string,
	stockQuantity int,
	createdAt time.Time,
) (Offer, error) {
	offer := Offer{
		id:            id,
		productID:     productID,
		regularPrice:  regularPrice,
		currentPrice:  currentPrice,
		currency:      currency,
		stockQuantity: stockQuantity,
		createdAt:     createdAt,
		updatedAt:     createdAt,
	}

	if err := offer.Validate(); err != nil {
		return Offer{}, err
	}

	return offer, nil
}

func (o Offer) Validate() error {
	if strings.TrimSpace(o.id) == "" {
		return ErrInvalidOfferID
	}

	if strings.TrimSpace(o.productID) == "" {
		return ErrInvalidProductID
	}

	if o.regularPrice <= 0 {
		return ErrInvalidRegularPrice
	}

	if o.currentPrice < 0 {
		return ErrInvalidCurrentPrice
	}

	if !isValidCurrency(o.currency) {
		return ErrInvalidCurrency
	}

	if o.stockQuantity < 0 {
		return ErrInvalidStockQuantity
	}

	if o.createdAt.IsZero() {
		return ErrInvalidCreatedAt
	}

	return nil
}

func isValidCurrency(currency string) bool {
	if len(currency) != CurrencyCodeLength {
		return false
	}

	for _, char := range currency {
		if char < 'A' || char > 'Z' {
			return false
		}
	}

	return true
}

func (o Offer) DiscountPercent() int {
	if o.regularPrice == 0 || o.currentPrice >= o.regularPrice {
		return 0
	}

	discount := (o.regularPrice - o.currentPrice) * 100 / o.regularPrice

	return int(discount)
}

func (o *Offer) UpdatePrice(newPrice int64, updatedAt time.Time) error {
	if newPrice < 0 {
		return ErrInvalidCurrentPrice
	}

	if updatedAt.IsZero() {
		return ErrInvalidUpdatedAt
	}

	o.currentPrice = newPrice
	o.updatedAt = updatedAt

	return nil
}

func (o *Offer) UpdateStock(newQuantity int, updatedAt time.Time) error {
	if newQuantity < 0 {
		return ErrInvalidStockQuantity
	}

	if updatedAt.IsZero() {
		return ErrInvalidUpdatedAt
	}

	o.stockQuantity = newQuantity
	o.updatedAt = updatedAt

	return nil
}

func (o Offer) ID() string {
	return o.id
}

func (o Offer) ProductID() string {
	return o.productID
}

func (o Offer) RegularPrice() int64 {
	return o.regularPrice
}

func (o Offer) CurrentPrice() int64 {
	return o.currentPrice
}

func (o Offer) Currency() string {
	return o.currency
}

func (o Offer) StockQuantity() int {
	return o.stockQuantity
}

func (o Offer) CreatedAt() time.Time {
	return o.createdAt
}

func (o Offer) UpdatedAt() time.Time {
	return o.updatedAt
}
