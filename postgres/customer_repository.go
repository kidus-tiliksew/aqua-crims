package postgres

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

func (c *CustomerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	res := &Customer{
		Name:  customer.Name,
		Email: customer.Email,
	}
	err := c.db.Create(res).Error
	if err != nil {
		return err
	}
	customer.ID = res.ID
	return nil
}

func (c *CustomerRepository) FindByID(ctx context.Context, id int64) (*domain.Customer, error) {
	res := &Customer{}
	err := c.db.First(res, id).Error
	if err != nil {
		return nil, err
	}
	return &domain.Customer{
		ID:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
