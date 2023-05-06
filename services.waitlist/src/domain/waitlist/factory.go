package waitlist

import (
	"github.com/google/uuid"
	"time"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newWaitlistErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) New(email string) *Entity {
	t := time.Now()
	return &Entity{
		Email:      email,
		LeaveToken: uuid.New().String(),
		IsActive:   true,
		UpdatedAt:  t,
		CreatedAt:  t,
	}
}
