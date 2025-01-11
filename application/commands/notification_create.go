package commands

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type NotificationCreate struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type NotificationCreateHandler struct {
	notifications domain.NotificationRepository
}

func NewNotificationCreateHandler(notifications domain.NotificationRepository) NotificationCreateHandler {
	return NotificationCreateHandler{
		notifications: notifications,
	}
}

func (h NotificationCreateHandler) NotificationCreate(ctx context.Context, cmd NotificationCreate) (*domain.Notification, error) {
	notification := domain.CreateNotification(cmd.UserID, cmd.Message)

	if err := h.notifications.Create(ctx, notification); err != nil {
		return nil, err
	}

	return notification, nil
}
