package domain

type Notification struct {
	ID      int64
	UserID  string
	Message string
}

type NotificationReceiver interface {
	SendStructuredMessage(userID, text string) error
}

func CreateNotification(userID string, message string) *Notification {
	return &Notification{
		UserID:  userID,
		Message: message,
	}
}
