package model

import (
	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/status"
)

type Product struct {
	BaseModel
	Name          string
	Description   string
	Price         float64
	Stock         int
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
