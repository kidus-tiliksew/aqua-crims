package domain

import "context"

type CloudResourceRepository interface {
	Create(context.Context, *CloudResource) error
	Update(context.Context, *CloudResource) error
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*CloudResource, error)
	FindByName(context.Context, string) (*CloudResource, error)
	FindByCustomer(context.Context, int64) ([]CloudResource, error)
}
