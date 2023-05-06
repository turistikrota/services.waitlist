package waitlist

import "time"

type Entity struct {
	UUID       string
	Email      string
	LeaveToken string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
