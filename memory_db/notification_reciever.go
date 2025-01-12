package memorydb

type InMemoryNotificationReceiver struct {
	messages []string
}

func (n *InMemoryNotificationReceiver) SendStructuredMessage(userID, message string) error {
	n.messages = append(n.messages, message)
	return nil
}
