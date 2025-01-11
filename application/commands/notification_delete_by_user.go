package commands

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type NotificationDeleteByUser struct {
	UserID string `json:"user_id"`
}

type NotificationDeleteByUserHandler struct {
	notifications domain.NotificationRepository
}

func NewNotificationDeleteByUserHandler(notifications domain.NotificationRepository) NotificationDeleteByUserHandler {
	return NotificationDeleteByUserHandler{
		notifications: notifications,
	}
}

func (h NotificationDeleteByUserHandler) NotificationDeleteByUser(ctx context.Context, userID string) error {
	return h.notifications.DeleteByUserID(ctx, userID)
}
