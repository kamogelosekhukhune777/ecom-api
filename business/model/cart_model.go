package model

import (
	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/quantity"
)

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
	Quantity  quantity.Quantity
	UnitPrice float64
}
