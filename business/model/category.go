package model

type Category struct {
	BaseModel
	Name        string
	Description string
	Products    []Product
}
