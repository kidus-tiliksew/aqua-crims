package commands

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceCreate struct {
	CustomerID int64  `json:"customer_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Region     string `json:"region"`
}

type CloudResourceCreateHandler struct {
	resources    domain.CloudResourceRepository
	notification domain.NotificationReceiver
}

func NewCloudResourceCreateHandler(resources domain.CloudResourceRepository, notification domain.NotificationReceiver) CloudResourceCreateHandler {
	return CloudResourceCreateHandler{
		resources:    resources,
		notification: notification,
	}
}

func (h CloudResourceCreateHandler) CloudResourceCreate(ctx context.Context, cmd CloudResourceCreate) (*domain.CloudResource, error) {
	resource, err := domain.CreateCloudResource(cmd.Name, cmd.Type, cmd.Region, cmd.CustomerID)
	if err != nil {
		return nil, err
	}

	// Create the resource
	if err := h.resources.Create(ctx, resource); err != nil {
		return nil, err
	}

	// Send event to the notification service
	if resource.CustomerID != 0 {
		if err := h.notification.SendStructuredMessage(fmt.Sprintf("%d", resource.CustomerID), fmt.Sprintf("New cloud resource %s created", resource.Name)); err != nil {
			slog.WarnContext(ctx, "failed to send notification", "Err", err)
		}
	}

	return resource, nil
}
