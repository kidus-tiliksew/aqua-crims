package postgres

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db}
}

var _ domain.NotificationRepository = (*NotificationRepository)(nil)

func (n *NotificationRepository) Create(ctx context.Context, notification *domain.Notification) error {
	res := &Notification{
		UserID:  notification.UserID,
		Message: notification.Message,
	}
	err := n.db.Create(res).Error
	if err != nil {
		return err
	}
	notification.ID = res.ID
	return nil
}

func (n *NotificationRepository) DeleteByID(ctx context.Context, id int64) error {
	return n.db.Where("id = ?", id).Delete(&Notification{}).Error
}

func (n *NotificationRepository) DeleteByUserID(ctx context.Context, id string) error {
	return n.db.Where("user_id = ?", id).Delete(&Notification{}).Error
}

func (n *NotificationRepository) FindByUserID(context.Context, string) ([]*domain.Notification, error) {
	var res []*Notification
	err := n.db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	var notifications []*domain.Notification
	for _, r := range res {
		notifications = append(notifications, &domain.Notification{
			ID:      r.ID,
			UserID:  r.UserID,
			Message: r.Message,
		})
	}
	return notifications, nil
}
