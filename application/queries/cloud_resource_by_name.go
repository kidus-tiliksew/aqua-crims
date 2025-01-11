package queries

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceGetByName struct {
	Name string `json:"name"`
}

type CloudResourceGetByNameHandler struct {
	resources domain.CloudResourceRepository
}

func NewCloudResourceGetByNameHandler(resources domain.CloudResourceRepository) CloudResourceGetByNameHandler {
	return CloudResourceGetByNameHandler{
		resources: resources,
	}
}

func (h CloudResourceGetByNameHandler) CloudResourceGetByName(ctx context.Context, query CloudResourceGetByName) (*domain.CloudResource, error) {
	return h.resources.FindByName(ctx, query.Name)
}