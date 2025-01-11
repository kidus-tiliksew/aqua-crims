package queries

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceGetByCustomer struct {
	CustomerID int64 `json:"customer_id"`
}

type CloudResourceGetByCustomerHandler struct {
	resources domain.CloudResourceRepository
}

func NewCloudResourceGetByCustomerHandler(resources domain.CloudResourceRepository) CloudResourceGetByCustomerHandler {
	return CloudResourceGetByCustomerHandler{
		resources: resources,
	}
}

func (h CloudResourceGetByCustomerHandler) CloudResourceGetByCustomer(ctx context.Context, query CloudResourceGetByCustomer) ([]domain.CloudResource, error) {
	return h.resources.FindByCustomer(ctx, query.CustomerID)
}
