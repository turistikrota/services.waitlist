package entity

import (
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"time"
)

type MongoWaitlist struct {
	UUID       string    `bson:"_id,omitempty"`
	Email      string    `bson:"email"`
	LeaveToken string    `bson:"leave_token"`
	IsActive   bool      `bson:"is_active"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

func (m *MongoWaitlist) FromEntity(e *waitlist.Entity) *MongoWaitlist {
	m.Email = e.Email
	m.LeaveToken = e.LeaveToken
	m.IsActive = e.IsActive
	m.CreatedAt = e.CreatedAt
	m.UpdatedAt = e.UpdatedAt
	return m
}

func (m *MongoWaitlist) ToEntity() *waitlist.Entity {
	return &waitlist.Entity{
		UUID:       m.UUID,
		Email:      m.Email,
		LeaveToken: m.LeaveToken,
		IsActive:   m.IsActive,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}
