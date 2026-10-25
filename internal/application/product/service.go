package product

import (
	"context"
	"time"

	"github.com/dmitdub/pricepulse/internal/domain"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
	ctx context.Context,
	id string,
	name string,
	description string,
	category string,
	createdAt time.Time,
) (domain.Product, error) {
	product, err := domain.NewProduct(
		id,
		name,
		description,
		category,
		createdAt,
	)
	if err != nil {
		return domain.Product{}, err
	}

	if err := s.repository.Save(ctx, product); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *Service) GetByID(
	ctx context.Context,
	id string,
) (domain.Product, error) {
	product, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
