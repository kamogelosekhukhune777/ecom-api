package model

import "github.com/google/uuid"

// Order represents an order in the system.
type Order struct {
	BaseModel
	UserId     uuid.UUID
	User       User
	OrderItems []OrderItem
	TotalPrice float64
	TotalItems int
	Status     string
	Address    string
}

// OrderItem represents an item in an order.
type OrderItem struct {
	BaseModel
	ProductId uuid.UUID
	Product   Product
	OrderId   uuid.UUID
	Order     Order
	Quantity  int
	UnitPrice float64
}
