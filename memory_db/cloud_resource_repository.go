package memorydb

import (
	"context"
	"errors"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type InMemoryCloudResourceRepository struct {
	Resources []domain.CloudResource
}

var ErrResourceNotFound = errors.New("resource not found")

var _ domain.CloudResourceRepository = (*InMemoryCloudResourceRepository)(nil)

// Delete implements domain.CloudResourceRepository.
func (r *InMemoryCloudResourceRepository) Delete(ctx context.Context, id int64) error {
	for i, resource := range r.Resources {
		if resource.ID == id {
			r.Resources = append(r.Resources[:i], r.Resources[i+1:]...)
			return nil
		}
	}
	return ErrResourceNotFound
}

// FindByCustomer implements domain.CloudResourceRepository.
func (r *InMemoryCloudResourceRepository) FindByCustomer(ctx context.Context, customerID int64) ([]domain.CloudResource, error) {
	var result []domain.CloudResource
	for _, resource := range r.Resources {
		if resource.CustomerID == customerID {
			result = append(result, resource)
		}
	}
	return result, nil
}

// FindByID implements domain.CloudResourceRepository.
func (r *InMemoryCloudResourceRepository) FindByID(ctx context.Context, id int64) (*domain.CloudResource, error) {
	for _, resource := range r.Resources {
		if resource.ID == id {
			return &resource, nil
		}
	}
	return nil, ErrResourceNotFound
}

// FindByName implements domain.CloudResourceRepository.
func (r *InMemoryCloudResourceRepository) FindByName(ctx context.Context, name string) (*domain.CloudResource, error) {
	for _, resource := range r.Resources {
		if resource.Name == name {
			return &resource, nil
		}
	}
	return nil, ErrResourceNotFound
}

// Update implements domain.CloudResourceRepository.
func (r *InMemoryCloudResourceRepository) Update(ctx context.Context, resource *domain.CloudResource) error {
	for i, res := range r.Resources {
		if res.ID == resource.ID {
			r.Resources[i] = *resource
			return nil
		}
	}
	return ErrResourceNotFound
}

func (r *InMemoryCloudResourceRepository) Create(ctx context.Context, resource *domain.CloudResource) error {
	r.Resources = append(r.Resources, *resource)
	return nil
}
