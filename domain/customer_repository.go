package domain

import "context"

type CustomerRepository interface {
	Create(context.Context, *Customer) error
	FindByID(context.Context, int64) (*Customer, error)
}
