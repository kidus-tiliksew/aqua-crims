package commands

import (
	"context"
	"fmt"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CustomerCreateCloudResources struct {
	CustomerID int64    `json:"customer_id"`
	Names      []string `json:"names"`
}

type CustomerCreateCloudResourcesHandler struct {
	resources domain.CloudResourceRepository
}

func NewCustomerCreateCloudResourcesHandler(resources domain.CloudResourceRepository) CustomerCreateCloudResourcesHandler {
	return CustomerCreateCloudResourcesHandler{
		resources: resources,
	}
}

func (h CustomerCreateCloudResourcesHandler) CustomerCreateCloudResources(ctx context.Context, cmd CustomerCreateCloudResources) ([]*domain.CloudResource, error) {
	var resources []*domain.CloudResource
	for _, name := range cmd.Names {
		// Check if the resource exists
		res, err := h.resources.FindByName(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("failed to find resource by name: %w", err)
		}

		res.CustomerID = cmd.CustomerID

		// Update the resource
		if err := h.resources.Update(ctx, res); err != nil {
			return nil, fmt.Errorf("failed to update resource: %w", err)
		}

		resources = append(resources, res)
	}

	return resources, nil
}
