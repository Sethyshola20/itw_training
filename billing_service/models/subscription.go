package models

import "time"

type Subscription struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubscriptionRequest struct {
	Subscription string `json:"subscription"`
}
