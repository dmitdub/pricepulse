package offer

import (
	"context"

	"github.com/dmitdub/pricepulse/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, offer domain.Offer) error
	FindByID(ctx context.Context, id string) (domain.Offer, error)
	FindByProductID(ctx context.Context, productID string) (domain.Offer, error)
}
