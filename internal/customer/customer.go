package customer

import "github.com/google/uuid"

type Customer struct {
	ID uuid.UUID
}

func New() *Customer {
	return &Customer{ID: uuid.New()}
}
