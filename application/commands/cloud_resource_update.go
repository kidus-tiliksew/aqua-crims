package commands

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CloudResourceUpdate struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Region     string `json:"region"`
	CustomerID int64  `json:"customer_id"`
}

type CloudResourceUpdateHandler struct {
	resources    domain.CloudResourceRepository
	customers    domain.CustomerRepository
	notification domain.NotificationReceiver
}

func NewCloudResourceUpdateHandler(resources domain.CloudResourceRepository, customers domain.CustomerRepository, notification domain.NotificationReceiver) CloudResourceUpdateHandler {
	return CloudResourceUpdateHandler{
		resources:    resources,
		customers:    customers,
		notification: notification,
	}
}

func (h CloudResourceUpdateHandler) CloudResourceUpdate(ctx context.Context, cmd CloudResourceUpdate) error {
	resource, err := domain.UpdateCloudResource(cmd.ID, cmd.CustomerID, cmd.Name, cmd.Type, cmd.Region)
	if err != nil {
		return err
	}

	// Check if the customer exists
	if _, err := h.customers.FindByID(ctx, cmd.CustomerID); err != nil {
		return fmt.Errorf("could not find customer with id %d", cmd.CustomerID)
	}

	// Check if the resource exists
	if _, err := h.resources.FindByID(ctx, cmd.ID); err != nil {
		return fmt.Errorf("could not find resource with id %d", cmd.ID)
	}

	// Send event to the notification service
	if err := h.notification.SendStructuredMessage(fmt.Sprintf("%d", cmd.CustomerID), fmt.Sprintf("Attached cloud resource %s", resource.Name)); err != nil {
		slog.WarnContext(ctx, "failed to send notification", "Err", err)
	}

	return h.resources.Update(ctx, resource)
}
