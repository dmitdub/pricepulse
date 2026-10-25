package offer

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
	productID string,
	regularPrice int64,
	currentPrice int64,
	currency string,
	stockQuantity int,
	createdAt time.Time,
) (domain.Offer, error) {
	offer, err := domain.NewOffer(
		id,
		productID,
		regularPrice,
		currentPrice,
		currency,
		stockQuantity,
		createdAt,
	)
	if err != nil {
		return domain.Offer{}, err
	}

	if err := s.repository.Save(ctx, offer); err != nil {
		return domain.Offer{}, err
	}

	return offer, nil
}

func (s *Service) GetByID(
	ctx context.Context,
	id string,
) (domain.Offer, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *Service) GetByProductID(
	ctx context.Context,
	productID string,
) (domain.Offer, error) {
	return s.repository.FindByProductID(ctx, productID)
}

func (s *Service) UpdatePrice(
	ctx context.Context,
	id string,
	newPrice int64,
	updatedAt time.Time,
) (domain.Offer, error) {
	offer, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return domain.Offer{}, err
	}

	if err := offer.UpdatePrice(newPrice, updatedAt); err != nil {
		return domain.Offer{}, err
	}

	if err := s.repository.Save(ctx, offer); err != nil {
		return domain.Offer{}, err
	}

	return offer, nil
}

func (s *Service) UpdateStock(
	ctx context.Context,
	id string,
	newQuantity int,
	updatedAt time.Time,
) (domain.Offer, error) {
	offer, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return domain.Offer{}, err
	}

	if err := offer.UpdateStock(newQuantity, updatedAt); err != nil {
		return domain.Offer{}, err
	}

	if err := s.repository.Save(ctx, offer); err != nil {
		return domain.Offer{}, err
	}

	return offer, nil
}
