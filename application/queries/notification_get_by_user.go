package queries

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type NotificationGetByUser struct {
	UserID string `json:"user_id"`
}

type NotificationGetByUserHandler struct {
	notifications domain.NotificationRepository
}

func NewNotificationGetByUserHandler(notifications domain.NotificationRepository) NotificationGetByUserHandler {
	return NotificationGetByUserHandler{
		notifications: notifications,
	}
}

func (h NotificationGetByUserHandler) NotificationGetByUser(ctx context.Context, userID string) ([]*domain.Notification, error) {
	return h.notifications.FindByUserID(ctx, userID)
}
