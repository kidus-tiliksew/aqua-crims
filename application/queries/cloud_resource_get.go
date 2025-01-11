package queries

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceGet struct {
	ID int64 `json:"id"`
}

type CloudResourceGetHandler struct {
	resources domain.CloudResourceRepository
}

func NewCloudResourceGetHandler(resources domain.CloudResourceRepository) CloudResourceGetHandler {
	return CloudResourceGetHandler{
		resources: resources,
	}
}

func (h CloudResourceGetHandler) CloudResourceGet(ctx context.Context, query CloudResourceGet) (*domain.CloudResource, error) {
	return h.resources.FindByID(ctx, query.ID)
}
