package model

import (
	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/quantity"
)

// Order represents an order in the system.
type Order struct {
	BaseModel
	UserId     uuid.UUID
	User       User
	OrderItems []OrderItem
	TotalPrice float64
	TotalItems quantity.Quantity
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
	Quantity  quantity.Quantity
	UnitPrice float64
}
