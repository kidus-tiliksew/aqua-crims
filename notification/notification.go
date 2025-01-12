package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationService struct {
	conn          *amqp.Connection
	channel       *amqp.Channel
	exchange      string
	notifications domain.NotificationRepository
}

type StructuredMessage struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

func NewNotificationService(notifications domain.NotificationRepository, dsn string, exchangeName string) (*NotificationService, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect RabbitMQ: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}
	if err := ch.ExchangeDeclare(
		exchangeName, "fanout", true, false, false, false, nil,
	); err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}
	ns := &NotificationService{
		conn: conn, channel: ch, exchange: exchangeName, notifications: notifications,
	}
	return ns, nil
}

func (ns *NotificationService) Subscribe() {
	q, err := ns.channel.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		log.Println("Queue declare error:", err)
		return
	}
	err = ns.channel.QueueBind(q.Name, "", ns.exchange, false, nil)
	if err != nil {
		log.Println("Queue bind error:", err)
		return
	}
	msgs, err := ns.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println("Consume error:", err)
		return
	}
	for d := range msgs {
		var sm StructuredMessage
		if err := json.Unmarshal(d.Body, &sm); err != nil {
			log.Println("Received notification:", string(d.Body))
		} else {
			log.Printf("Received message from user %s: %s\n", sm.UserID, sm.Text)
			ns.notifications.Create(context.Background(), &domain.Notification{UserID: sm.UserID, Message: sm.Text})
		}
	}
}

func (ns *NotificationService) SendStructuredMessage(userID, text string) error {
	sm := StructuredMessage{UserID: userID, Text: text}
	body, err := json.Marshal(sm)
	if err != nil {
		return err
	}
	return ns.channel.Publish(ns.exchange, "", false, false,
		amqp.Publishing{ContentType: "application/json", Body: body},
	)
}
