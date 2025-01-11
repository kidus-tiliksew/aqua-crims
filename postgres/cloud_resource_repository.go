package postgres

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"gorm.io/gorm"
)

type CloudResourceRepository struct {
	db *gorm.DB
}

func NewCloudResourceRepository(db *gorm.DB) *CloudResourceRepository {
	return &CloudResourceRepository{
		db: db,
	}
}

var _ domain.CloudResourceRepository = (*CloudResourceRepository)(nil)

func (c *CloudResourceRepository) Create(ctx context.Context, resource *domain.CloudResource) error {
	res := &CloudResource{
		CustomerID: resource.CustomerID,
		Name:       resource.Name,
		Type:       resource.Type,
		Region:     resource.Region,
	}
	err := c.db.Create(res).Error
	if err != nil {
		return err
	}

	resource.ID = res.ID
	return nil
}

func (c *CloudResourceRepository) Delete(ctx context.Context, id int64) error {
	return c.db.Delete(&CloudResource{}, id).Error
}

func (c *CloudResourceRepository) FindByID(ctx context.Context, id int64) (*domain.CloudResource, error) {
	var res CloudResource
	err := c.db.Where("id = ?", id).First(&res).Error
	if err != nil {
		return nil, err
	}

	return &domain.CloudResource{
		ID:         res.ID,
		Name:       res.Name,
		Type:       res.Type,
		Region:     res.Region,
		CustomerID: res.CustomerID,
	}, nil
}

func (c *CloudResourceRepository) FindByName(ctx context.Context, name string) (*domain.CloudResource, error) {
	var res CloudResource
	err := c.db.Where("name = ?", name).First(&res).Error
	if err != nil {
		return nil, err
	}

	return &domain.CloudResource{
		ID:         res.ID,
		Name:       res.Name,
		Type:       res.Type,
		Region:     res.Region,
		CustomerID: res.CustomerID,
	}, nil
}

func (c *CloudResourceRepository) FindByCustomer(ctx context.Context, id int64) ([]domain.CloudResource, error) {
	var res []*CloudResource
	err := c.db.Where("customer_id = ?", id).Find(&res).Error
	if err != nil {
		return nil, err
	}

	var resources []domain.CloudResource
	for _, r := range res {
		resources = append(resources, domain.CloudResource{
			ID:         r.ID,
			Name:       r.Name,
			Type:       r.Type,
			Region:     r.Region,
			CustomerID: r.CustomerID,
		})
	}

	return resources, nil
}

func (c *CloudResourceRepository) Update(ctx context.Context, resource *domain.CloudResource) error {
	return c.db.Model(&CloudResource{}).Where("id = ?", resource.ID).Updates(CloudResource{
		Name:       resource.Name,
		Type:       resource.Type,
		Region:     resource.Region,
		CustomerID: resource.CustomerID,
	}).Error
}
