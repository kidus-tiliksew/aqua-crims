package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"github.com/stretchr/testify/assert"
)

type inMemoryCloudResourceRepository struct {
	resources []domain.CloudResource
}

var ErrResourceNotFound = errors.New("resource not found")

// Delete implements domain.CloudResourceRepository.
func (r *inMemoryCloudResourceRepository) Delete(ctx context.Context, id int64) error {
	for i, resource := range r.resources {
		if resource.ID == id {
			r.resources = append(r.resources[:i], r.resources[i+1:]...)
			return nil
		}
	}
	return ErrResourceNotFound
}

// FindByCustomer implements domain.CloudResourceRepository.
func (r *inMemoryCloudResourceRepository) FindByCustomer(ctx context.Context, customerID int64) ([]domain.CloudResource, error) {
	var result []domain.CloudResource
	for _, resource := range r.resources {
		if resource.CustomerID == customerID {
			result = append(result, resource)
		}
	}
	return result, nil
}

// FindByID implements domain.CloudResourceRepository.
func (r *inMemoryCloudResourceRepository) FindByID(ctx context.Context, id int64) (*domain.CloudResource, error) {
	for _, resource := range r.resources {
		if resource.ID == id {
			return &resource, nil
		}
	}
	return nil, ErrResourceNotFound
}

// FindByName implements domain.CloudResourceRepository.
func (r *inMemoryCloudResourceRepository) FindByName(ctx context.Context, name string) (*domain.CloudResource, error) {
	for _, resource := range r.resources {
		if resource.Name == name {
			return &resource, nil
		}
	}
	return nil, ErrResourceNotFound
}

// Update implements domain.CloudResourceRepository.
func (r *inMemoryCloudResourceRepository) Update(ctx context.Context, resource *domain.CloudResource) error {
	for i, res := range r.resources {
		if res.ID == resource.ID {
			r.resources[i] = *resource
			return nil
		}
	}
	return ErrResourceNotFound
}

func (r *inMemoryCloudResourceRepository) Create(ctx context.Context, resource *domain.CloudResource) error {
	r.resources = append(r.resources, *resource)
	return nil
}



func TestInMemoryCloudResourceRepository_Create(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{}
	resource := &domain.CloudResource{
		ID:         1,
		CustomerID: 1,
		Name:       "Test Resource",
		Type:       "S3",
		Region:     "us-west-2",
	}

	err := repo.Create(context.Background(), resource)
	assert.NoError(t, err)
	assert.Len(t, repo.resources, 1)
	assert.Equal(t, resource, &repo.resources[0])
}

func TestInMemoryCloudResourceRepository_FindByID(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{
		resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}

	resource, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, resource)
	assert.Equal(t, int64(1), resource.ID)

	resource, err = repo.FindByID(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestInMemoryCloudResourceRepository_FindByName(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{
		resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}

	resource, err := repo.FindByName(context.Background(), "Test Resource")
	assert.NoError(t, err)
	assert.NotNil(t, resource)
	assert.Equal(t, "Test Resource", resource.Name)

	resource, err = repo.FindByName(context.Background(), "Nonexistent Resource")
	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestInMemoryCloudResourceRepository_FindByCustomer(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{
		resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource 1", Type: "S3", Region: "us-west-2"},
			{ID: 2, CustomerID: 1, Name: "Test Resource 2", Type: "EC2", Region: "us-east-1"},
			{ID: 3, CustomerID: 2, Name: "Test Resource 3", Type: "RDS", Region: "us-west-1"},
		},
	}

	resources, err := repo.FindByCustomer(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, resources, 2)

	resources, err = repo.FindByCustomer(context.Background(), 2)
	assert.NoError(t, err)
	assert.Len(t, resources, 1)

	resources, err = repo.FindByCustomer(context.Background(), 3)
	assert.NoError(t, err)
	assert.Len(t, resources, 0)
}

func TestInMemoryCloudResourceRepository_Update(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{
		resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}

	updatedResource := &domain.CloudResource{
		ID:         1,
		CustomerID: 1,
		Name:       "Updated Resource",
		Type:       "S3",
		Region:     "us-west-2",
	}

	err := repo.Update(context.Background(), updatedResource)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Resource", repo.resources[0].Name)

	nonexistentResource := &domain.CloudResource{
		ID:         2,
		CustomerID: 1,
		Name:       "Nonexistent Resource",
		Type:       "S3",
		Region:     "us-west-2",
	}

	err = repo.Update(context.Background(), nonexistentResource)
	assert.Error(t, err)
}

func TestInMemoryCloudResourceRepository_Delete(t *testing.T) {
	repo := &inMemoryCloudResourceRepository{
		resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}

	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, repo.resources, 0)

	err = repo.Delete(context.Background(), 2)
	assert.Error(t, err)
}
