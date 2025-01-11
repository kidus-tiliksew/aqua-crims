package domain

import "context"

type NotificationRepository interface {
	Create(context.Context, *Notification) error
	DeleteByID(context.Context, int64) error
	FindByUserID(context.Context, string) ([]*Notification, error)
	DeleteByUserID(context.Context, string) error
}
