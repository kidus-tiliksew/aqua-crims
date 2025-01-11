package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"github.com/stretchr/testify/assert"
)

type inMemoryCustomerRepository struct {
	customers []domain.Customer
}

var ErrCustomerNotFound = errors.New("customer not found")

// Create implements domain.CustomerRepository.
func (r *inMemoryCustomerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	r.customers = append(r.customers, *customer)
	return nil
}

// FindByID implements domain.CustomerRepository.
func (r *inMemoryCustomerRepository) FindByID(ctx context.Context, id int64) (*domain.Customer, error) {
	for _, customer := range r.customers {
		if customer.ID == id {
			return &customer, nil
		}
	}
	return nil, ErrCustomerNotFound
}

// FindByName implements domain.CustomerRepository.
func (r *inMemoryCustomerRepository) FindByName(ctx context.Context, name string) (*domain.Customer, error) {
	for _, customer := range r.customers {
		if customer.Name == name {
			return &customer, nil
		}
	}
	return nil, ErrCustomerNotFound
}

// Update implements domain.CustomerRepository.
func (r *inMemoryCustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	for i, cust := range r.customers {
		if cust.ID == customer.ID {
			r.customers[i] = *customer
			return nil
		}
	}
	return ErrCustomerNotFound
}

// Delete implements domain.CustomerRepository.
func (r *inMemoryCustomerRepository) Delete(ctx context.Context, id int64) error {
	for i, customer := range r.customers {
		if customer.ID == id {
			r.customers = append(r.customers[:i], r.customers[i+1:]...)
			return nil
		}
	}
	return ErrCustomerNotFound
}

func TestInMemoryCustomerRepository_Create(t *testing.T) {
	repo := &inMemoryCustomerRepository{}
	customer := &domain.Customer{
		ID:   1,
		Name: "Test Customer",
	}

	err := repo.Create(context.Background(), customer)
	assert.NoError(t, err)
	assert.Len(t, repo.customers, 1)
	assert.Equal(t, customer, &repo.customers[0])
}

func TestInMemoryCustomerRepository_FindByID(t *testing.T) {
	repo := &inMemoryCustomerRepository{
		customers: []domain.Customer{
			{ID: 1, Name: "Test Customer"},
		},
	}

	customer, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, int64(1), customer.ID)

	customer, err = repo.FindByID(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestInMemoryCustomerRepository_FindByName(t *testing.T) {
	repo := &inMemoryCustomerRepository{
		customers: []domain.Customer{
			{ID: 1, Name: "Test Customer"},
		},
	}

	customer, err := repo.FindByName(context.Background(), "Test Customer")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Test Customer", customer.Name)

	customer, err = repo.FindByName(context.Background(), "Nonexistent Customer")
	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestInMemoryCustomerRepository_Update(t *testing.T) {
	repo := &inMemoryCustomerRepository{
		customers: []domain.Customer{
			{ID: 1, Name: "Test Customer"},
		},
	}

	updatedCustomer := &domain.Customer{
		ID:   1,
		Name: "Updated Customer",
	}

	err := repo.Update(context.Background(), updatedCustomer)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Customer", repo.customers[0].Name)

	nonexistentCustomer := &domain.Customer{
		ID:   2,
		Name: "Nonexistent Customer",
	}

	err = repo.Update(context.Background(), nonexistentCustomer)
	assert.Error(t, err)
}

func TestInMemoryCustomerRepository_Delete(t *testing.T) {
	repo := &inMemoryCustomerRepository{
		customers: []domain.Customer{
			{ID: 1, Name: "Test Customer"},
		},
	}

	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, repo.customers, 0)

	err = repo.Delete(context.Background(), 2)
	assert.Error(t, err)
}
