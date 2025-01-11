package domain

import (
	"net/mail"

	"github.com/stackus/errors"
)

var (
	ErrCustomerNameEmpty      = errors.Wrap(errors.ErrBadRequest, "name cannot be blank")
	ErrCustomerEmailEmpty     = errors.Wrap(errors.ErrBadRequest, "email cannot be blank")
	ErrCustomerEmailInvalid   = errors.Wrap(errors.ErrConflict, "email is invalid")
	ErrCustomerEmailNotUnique = errors.Wrap(errors.ErrConflict, "email must be unique")
)

type Customer struct {
	ID    int64
	Name  string
	Email string
}

func CreateCustomer(name string, email string) (*Customer, error) {
	if name == "" {
		return nil, ErrCustomerNameEmpty
	}

	if email == "" {
		return nil, ErrCustomerEmailEmpty
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return nil, ErrCustomerEmailInvalid
	}

	return &Customer{
		Name:  name,
		Email: email,
	}, nil
}
