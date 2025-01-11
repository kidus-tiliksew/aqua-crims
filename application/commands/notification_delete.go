package commands

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type NotificationDelete struct {
	ID int64 `json:"id"`
}

type NotificationDeleteHandler struct {
	notifications domain.NotificationRepository
}

func NewNotificationDeleteHandler(notifications domain.NotificationRepository) NotificationDeleteHandler {
	return NotificationDeleteHandler{
		notifications: notifications,
	}
}

func (h NotificationDeleteHandler) NotificationDelete(ctx context.Context, id int64) error {
	return h.notifications.DeleteByID(ctx, id)
}
