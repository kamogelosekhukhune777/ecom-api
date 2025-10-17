package model

import "github.com/google/uuid"

// Cart represents a shopping cart in the system.
type Cart struct {
	BaseModel
	UserId    uuid.UUID
	User      User
	CartItems []CartItem
}

// CartItem represents an item in the shopping cart.
type CartItem struct {
	BaseModel
	ProductId uuid.UUID
	Product   Product
	CartId    uuid.UUID
	Cart      Cart
	Quantity  int
	UnitPrice float64
}
