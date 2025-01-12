package memorydb

import (
	"context"
	"errors"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type InMemoryNotificationRepository struct {
	Notifications []domain.Notification
}

var ErrNotificationNotFound = errors.New("notification not found")

func NewInMemoryNotificationRepository() *InMemoryNotificationRepository {
	return &InMemoryNotificationRepository{}
}

var _ domain.NotificationRepository = (*InMemoryNotificationRepository)(nil)

func (r *InMemoryNotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	r.Notifications = append(r.Notifications, *notification)
	return nil
}

func (r *InMemoryNotificationRepository) FindByID(ctx context.Context, id int64) (*domain.Notification, error) {
	for _, notification := range r.Notifications {
		if notification.ID == id {
			return &notification, nil
		}
	}
	return nil, ErrNotificationNotFound
}

func (r *InMemoryNotificationRepository) FindByUserID(ctx context.Context, userID string) ([]*domain.Notification, error) {
	var result []*domain.Notification
	for i := range r.Notifications {
		if r.Notifications[i].UserID == userID {
			result = append(result, &r.Notifications[i])
		}
	}
	return result, nil
}

func (r *InMemoryNotificationRepository) Delete(ctx context.Context, id int64) error {
	for i, notification := range r.Notifications {
		if notification.ID == id {
			r.Notifications = append(r.Notifications[:i], r.Notifications[i+1:]...)
			return nil
		}
	}
	return ErrNotificationNotFound
}

func (r *InMemoryNotificationRepository) DeleteByID(ctx context.Context, id int64) error {
	for i, notification := range r.Notifications {
		if notification.ID == id {
			r.Notifications = append(r.Notifications[:i], r.Notifications[i+1:]...)
			return nil
		}
	}
	return ErrNotificationNotFound
}

func (r *InMemoryNotificationRepository) DeleteByUserID(ctx context.Context, userID string) error {
	var remaining []domain.Notification
	for _, notification := range r.Notifications {
		if notification.UserID != userID {
			remaining = append(remaining, notification)
		}
	}
	r.Notifications = remaining
	return nil
}
