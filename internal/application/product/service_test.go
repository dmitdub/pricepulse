package product

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dmitdub/pricepulse/internal/domain"
)

type mockRepository struct {
	saveErr      error
	product      domain.Product
	findErr      error
	savedProduct domain.Product
	saveCalls    int
}

func (m *mockRepository) Save(
	ctx context.Context,
	product domain.Product,
) error {
	m.saveCalls++
	m.savedProduct = product

	return m.saveErr
}

func (m *mockRepository) FindByID(
	ctx context.Context,
	id string,
) (domain.Product, error) {
	if m.findErr != nil {
		return domain.Product{}, m.findErr
	}

	return m.product, nil
}

func TestServiceCreate(t *testing.T) {
	createdAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)

	repository := &mockRepository{
		saveErr:      nil,
		product:      domain.Product{},
		findErr:      nil,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		"product-1",
		"Smartphone",
		"Modern smartphone",
		"Electronics",
		createdAt,
	)
	if err != nil {
		t.Fatalf("Create() error = %v, want nil", err)
	}

	if product.ID() != "product-1" {
		t.Errorf("ID() = %q, want %q", product.ID(), "product-1")
	}

	if repository.saveCalls != 1 {
		t.Errorf("Save() calls = %d, want 1", repository.saveCalls)
	}

	if repository.savedProduct.ID() != product.ID() {
		t.Errorf(
			"saved product ID = %q, want %q",
			repository.savedProduct.ID(),
			product.ID(),
		)
	}
}

func TestServiceCreateInvalidProduct(t *testing.T) {
	repository := &mockRepository{
		saveErr:      nil,
		product:      domain.Product{},
		findErr:      nil,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		"",
		"Smartphone",
		"Modern smartphone",
		"Electronics",
		time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC),
	)

	if !errors.Is(err, domain.ErrInvalidProductID) {
		t.Fatalf("Create() error = %v, want %v", err, domain.ErrInvalidProductID)
	}

	if product.ID() != "" {
		t.Errorf("ID() = %q, want empty string", product.ID())
	}

	if repository.saveCalls != 0 {
		t.Errorf("Save() calls = %d, want 0", repository.saveCalls)
	}
}

func TestServiceCreateRepositoryError(t *testing.T) {
	repositoryErr := errors.New("database unavailable")

	repository := &mockRepository{
		saveErr:      repositoryErr,
		product:      domain.Product{},
		findErr:      nil,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.Create(
		context.Background(),
		"product-1",
		"Smartphone",
		"Modern smartphone",
		"Electronics",
		time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC),
	)

	if !errors.Is(err, repositoryErr) {
		t.Fatalf("Create() error = %v, want %v", err, repositoryErr)
	}

	if product.ID() != "" {
		t.Errorf("ID() = %q, want empty string", product.ID())
	}
}

func TestServiceGetByID(t *testing.T) {
	expectedProduct, err := domain.NewProduct(
		"product-1",
		"Smartphone",
		"Modern smartphone",
		"Electronics",
		time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatalf("NewProduct() error = %v", err)
	}

	repository := &mockRepository{
		saveErr:      nil,
		product:      expectedProduct,
		findErr:      nil,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.GetByID(context.Background(), "product-1")
	if err != nil {
		t.Fatalf("GetByID() error = %v, want nil", err)
	}

	if product.ID() != expectedProduct.ID() {
		t.Errorf("ID() = %q, want %q", product.ID(), expectedProduct.ID())
	}
}

func TestServiceGetByIDRepositoryError(t *testing.T) {
	repositoryErr := errors.New("database unavailable")

	repository := &mockRepository{
		saveErr:      nil,
		product:      domain.Product{},
		findErr:      repositoryErr,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.GetByID(context.Background(), "product-1")

	if !errors.Is(err, repositoryErr) {
		t.Fatalf("GetByID() error = %v, want %v", err, repositoryErr)
	}

	if product.ID() != "" {
		t.Errorf("ID() = %q, want empty string", product.ID())
	}
}

func TestServiceGetByIDNotFound(t *testing.T) {
	repository := &mockRepository{
		saveErr:      nil,
		product:      domain.Product{},
		findErr:      domain.ErrProductNotFound,
		savedProduct: domain.Product{},
		saveCalls:    0,
	}

	service := NewService(repository)

	product, err := service.GetByID(context.Background(), "product-1")

	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Fatalf("GetByID() error = %v, want %v", err, domain.ErrProductNotFound)
	}

	if product.ID() != "" {
		t.Errorf("ID() = %q, want empty string", product.ID())
	}
}
