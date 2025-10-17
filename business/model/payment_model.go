package model

import "github.com/google/uuid"

// Payment represents a payment in the system.
type Payment struct {
	BaseModel
	Amount      float64
	Status      string
	AuthorityId string
	RefId       int
	UserId      uuid.UUID
	User        User
	OrderId     uuid.UUID
	Order       Order
}
