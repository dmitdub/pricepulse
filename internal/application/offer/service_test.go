package offer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dmitdub/pricepulse/internal/domain"
)

type mockRepository struct {
	saveErr       error
	offer         domain.Offer
	findErr       error
	savedOffer    domain.Offer
	saveCalls     int
	findByIDCalls int
}

func (m *mockRepository) Save(
	ctx context.Context,
	offer domain.Offer,
) error {
	m.saveCalls++
	m.savedOffer = offer

	return m.saveErr
}

func (m *mockRepository) FindByID(
	ctx context.Context,
	id string,
) (domain.Offer, error) {
	m.findByIDCalls++

	if m.findErr != nil {
		return domain.Offer{}, m.findErr
	}

	return m.offer, nil
}

func (m *mockRepository) FindByProductID(
	ctx context.Context,
	productID string,
) (domain.Offer, error) {
	if m.findErr != nil {
		return domain.Offer{}, m.findErr
	}

	return m.offer, nil
}

func newTestOffer(t *testing.T) domain.Offer {
	t.Helper()

	offer, err := domain.NewOffer(
		"offer-1",
		"product-1",
		50000,
		45000,
		"EUR",
		10,
		time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatalf("NewOffer() error = %v", err)
	}

	return offer
}

func TestServiceCreate(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)

	repository := &mockRepository{
		saveErr:       nil,
		offer:         domain.Offer{},
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.Create(
		context.Background(),
		"offer-1",
		"product-1",
		50000,
		45000,
		"EUR",
		10,
		createdAt,
	)
	if err != nil {
		t.Fatalf("Create() error = %v, want nil", err)
	}

	if offer.ID() != "offer-1" {
		t.Errorf("ID() = %q, want %q", offer.ID(), "offer-1")
	}

	if repository.saveCalls != 1 {
		t.Errorf("Save() calls = %d, want 1", repository.saveCalls)
	}

	if repository.savedOffer.ID() != offer.ID() {
		t.Errorf("saved offer ID = %q, want %q", repository.savedOffer.ID(), offer.ID())
	}
}

func TestServiceCreateInvalidOffer(t *testing.T) {
	repository := &mockRepository{
		saveErr:       nil,
		offer:         domain.Offer{},
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.Create(
		context.Background(),
		"",
		"product-1",
		50000,
		45000,
		"EUR",
		10,
		time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC),
	)

	if !errors.Is(err, domain.ErrInvalidOfferID) {
		t.Fatalf("Create() error = %v, want %v", err, domain.ErrInvalidOfferID)
	}

	if offer.ID() != "" {
		t.Errorf("ID() = %q, want empty string", offer.ID())
	}

	if repository.saveCalls != 0 {
		t.Errorf("Save() calls = %d, want 0", repository.saveCalls)
	}
}

func TestServiceGetByID(t *testing.T) {
	expectedOffer := newTestOffer(t)

	repository := &mockRepository{
		saveErr:       nil,
		offer:         expectedOffer,
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.GetByID(context.Background(), "offer-1")
	if err != nil {
		t.Fatalf("GetByID() error = %v, want nil", err)
	}

	if offer.ID() != expectedOffer.ID() {
		t.Errorf("ID() = %q, want %q", offer.ID(), expectedOffer.ID())
	}

	if repository.findByIDCalls != 1 {
		t.Errorf("FindByID() calls = %d, want 1", repository.findByIDCalls)
	}
}

func TestServiceGetByProductID(t *testing.T) {
	expectedOffer := newTestOffer(t)

	repository := &mockRepository{
		saveErr:       nil,
		offer:         expectedOffer,
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.GetByProductID(context.Background(), "product-1")
	if err != nil {
		t.Fatalf("GetByProductID() error = %v, want nil", err)
	}

	if offer.ProductID() != expectedOffer.ProductID() {
		t.Errorf("ProductID() = %q, want %q", offer.ProductID(), expectedOffer.ProductID())
	}
}

func TestServiceUpdatePrice(t *testing.T) {
	expectedOffer := newTestOffer(t)
	updatedAt := time.Date(2026, time.January, 2, 12, 0, 0, 0, time.UTC)

	repository := &mockRepository{
		saveErr:       nil,
		offer:         expectedOffer,
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.UpdatePrice(
		context.Background(),
		"offer-1",
		40000,
		updatedAt,
	)
	if err != nil {
		t.Fatalf("UpdatePrice() error = %v, want nil", err)
	}

	if offer.CurrentPrice() != 40000 {
		t.Errorf("CurrentPrice() = %d, want 40000", offer.CurrentPrice())
	}

	if !offer.UpdatedAt().Equal(updatedAt) {
		t.Errorf("UpdatedAt() = %v, want %v", offer.UpdatedAt(), updatedAt)
	}

	if repository.saveCalls != 1 {
		t.Errorf("Save() calls = %d, want 1", repository.saveCalls)
	}

	if repository.savedOffer.CurrentPrice() != 40000 {
		t.Errorf("saved CurrentPrice() = %d, want 40000", repository.savedOffer.CurrentPrice())
	}
}

func TestServiceUpdateStock(t *testing.T) {
	expectedOffer := newTestOffer(t)
	updatedAt := time.Date(2026, time.January, 2, 12, 0, 0, 0, time.UTC)

	repository := &mockRepository{
		saveErr:       nil,
		offer:         expectedOffer,
		findErr:       nil,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.UpdateStock(
		context.Background(),
		"offer-1",
		5,
		updatedAt,
	)
	if err != nil {
		t.Fatalf("UpdateStock() error = %v, want nil", err)
	}

	if offer.StockQuantity() != 5 {
		t.Errorf("StockQuantity() = %d, want 5", offer.StockQuantity())
	}

	if !offer.UpdatedAt().Equal(updatedAt) {
		t.Errorf("UpdatedAt() = %v, want %v", offer.UpdatedAt(), updatedAt)
	}

	if repository.saveCalls != 1 {
		t.Errorf("Save() calls = %d, want 1", repository.saveCalls)
	}

	if repository.savedOffer.StockQuantity() != 5 {
		t.Errorf("saved StockQuantity() = %d, want 5", repository.savedOffer.StockQuantity())
	}
}

func TestServiceUpdatePriceRepositoryError(t *testing.T) {
	repositoryErr := errors.New("database unavailable")

	repository := &mockRepository{
		saveErr:       nil,
		offer:         domain.Offer{},
		findErr:       repositoryErr,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.UpdatePrice(
		context.Background(),
		"offer-1",
		40000,
		time.Date(2026, time.January, 2, 12, 0, 0, 0, time.UTC),
	)

	if !errors.Is(err, repositoryErr) {
		t.Fatalf("UpdatePrice() error = %v, want %v", err, repositoryErr)
	}

	if offer.ID() != "" {
		t.Errorf("ID() = %q, want empty string", offer.ID())
	}

	if repository.saveCalls != 0 {
		t.Errorf("Save() calls = %d, want 0", repository.saveCalls)
	}
}

func TestServiceGetByIDNotFound(t *testing.T) {
	repository := &mockRepository{
		saveErr:       nil,
		offer:         domain.Offer{},
		findErr:       domain.ErrOfferNotFound,
		savedOffer:    domain.Offer{},
		saveCalls:     0,
		findByIDCalls: 0,
	}

	service := NewService(repository)

	offer, err := service.GetByID(context.Background(), "offer-1")

	if !errors.Is(err, domain.ErrOfferNotFound) {
		t.Fatalf("GetByID() error = %v, want %v", err, domain.ErrOfferNotFound)
	}

	if offer.ID() != "" {
		t.Errorf("ID() = %q, want empty string", offer.ID())
	}
}
