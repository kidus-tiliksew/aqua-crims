package memorydb

import (
	"context"
	"errors"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type InMemoryCustomerRepository struct {
	Customers []domain.Customer
}

var ErrCustomerNotFound = errors.New("customer not found")

var _ domain.CustomerRepository = (*InMemoryCustomerRepository)(nil)

// Create implements domain.CustomerRepository.
func (r *InMemoryCustomerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	r.Customers = append(r.Customers, *customer)
	return nil
}

// FindByID implements domain.CustomerRepository.
func (r *InMemoryCustomerRepository) FindByID(ctx context.Context, id int64) (*domain.Customer, error) {
	for _, customer := range r.Customers {
		if customer.ID == id {
			return &customer, nil
		}
	}
	return nil, ErrCustomerNotFound
}

// FindByName implements domain.CustomerRepository.
func (r *InMemoryCustomerRepository) FindByName(ctx context.Context, name string) (*domain.Customer, error) {
	for _, customer := range r.Customers {
		if customer.Name == name {
			return &customer, nil
		}
	}
	return nil, ErrCustomerNotFound
}

// Update implements domain.CustomerRepository.
func (r *InMemoryCustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	for i, cust := range r.Customers {
		if cust.ID == customer.ID {
			r.Customers[i] = *customer
			return nil
		}
	}
	return ErrCustomerNotFound
}

// Delete implements domain.CustomerRepository.
func (r *InMemoryCustomerRepository) Delete(ctx context.Context, id int64) error {
	for i, customer := range r.Customers {
		if customer.ID == id {
			r.Customers = append(r.Customers[:i], r.Customers[i+1:]...)
			return nil
		}
	}
	return ErrCustomerNotFound
}
