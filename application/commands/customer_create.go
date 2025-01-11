package commands

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type CustomerCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerCreateHandler struct {
	customers    domain.CustomerRepository
	notification domain.NotificationReceiver
}

func NewCustomerCreateHandler(customers domain.CustomerRepository, notification domain.NotificationReceiver) CustomerCreateHandler {
	return CustomerCreateHandler{
		customers:    customers,
		notification: notification,
	}
}

func (h CustomerCreateHandler) CustomerCreate(ctx context.Context, cmd CustomerCreate) (*domain.Customer, error) {
	customer, err := domain.CreateCustomer(cmd.Name, cmd.Email)
	if err != nil {
		return nil, err
	}

	if err := h.customers.Create(ctx, customer); err != nil {
		return nil, err
	}

	if err := h.notification.SendStructuredMessage(fmt.Sprintf("%d", customer.ID), "Account created"); err != nil {
		slog.WarnContext(ctx, "failed to send notification", "Err", err)
	}

	return customer, nil
}
