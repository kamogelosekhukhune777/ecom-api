package productbus

import (
	"time"

	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/quantity"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/role"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/status"
)

type BaseModel struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Product struct {
	BaseModel
	Name          string
	Description   string
	Price         float64
	Stock         quantity.Quantity
	CategoryID    uuid.UUID
	Status        status.Status
	Slug          string
	Category      Category
	AverageRating float64
	CountViews    int
	Images        []ProductImage
	Reviews       []ProductReview
}

type ProductImage struct {
	BaseModel
	ProductID uuid.UUID
	Product   Product
	Image     File
	ImageID   uuid.UUID
	IsMain    bool
}

type ProductReview struct {
	BaseModel
	ProductID uuid.UUID
	Rating    int
	Comment   string
	Product   Product
	UserId    uuid.UUID
	User      User
}

//====================================================================================

type NewProduct struct{}

type UpdateProduct struct{}

//====================================================================================

type File struct {
	BaseModel
	Name        string
	Directory   string
	Description string
	MimeType    string
}

type User struct {
	BaseModel
	Username     string
	FirstName    string
	LastName     string
	MobileNumber string
	Email        string
	Password     string
	Enabled      bool
	UserRoles    []role.Role
}

type Category struct {
	BaseModel
	Name        string
	Description string
	Products    []Product
}
