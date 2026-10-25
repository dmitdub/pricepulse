package product

import (
	"context"

	"github.com/dmitdub/pricepulse/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, product domain.Product) error
	FindByID(ctx context.Context, id string) (domain.Product, error)
}
