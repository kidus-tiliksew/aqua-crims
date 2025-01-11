package commands

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceDelete struct {
	ID int64 `json:"id"`
}

type CloudResourceDeleteHandler struct {
	resources domain.CloudResourceRepository
}

func NewCloudResourceDeleteHandler(resources domain.CloudResourceRepository) CloudResourceDeleteHandler {
	return CloudResourceDeleteHandler{
		resources: resources,
	}
}

func (h CloudResourceDeleteHandler) CloudResourceDelete(ctx context.Context, cmd CloudResourceDelete) error {
	return h.resources.Delete(ctx, cmd.ID)
}
