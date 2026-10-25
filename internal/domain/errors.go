package domain

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrOfferNotFound   = errors.New("offer not found")

	ErrInvalidOfferID            = errors.New("invalid offer ID")
	ErrInvalidProductID          = errors.New("invalid product ID")
	ErrInvalidProductName        = errors.New("product name cannot be empty")
	ErrInvalidProductDescription = errors.New("product description cannot be empty")
	ErrInvalidProductCategory    = errors.New("product category cannot be empty")

	ErrInvalidRegularPrice  = errors.New("regular price must be greater than zero")
	ErrInvalidCurrentPrice  = errors.New("current price cannot be negative")
	ErrInvalidCurrency      = errors.New("currency must be a 3-letter uppercase code")
	ErrInvalidStockQuantity = errors.New("stock quantity cannot be negative")

	ErrInvalidCreatedAt        = errors.New("created at cannot be zero")
	ErrInvalidUpdatedAt        = errors.New("updated at cannot be zero")
	ErrInvalidProductCreatedAt = errors.New("product created at cannot be zero")
	ErrInvalidProductUpdatedAt = errors.New("product updated at cannot be zero")
)
