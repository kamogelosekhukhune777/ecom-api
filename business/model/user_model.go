package model

import "github.com/kamogelosekhukhune777/ecom-api/business/types/role"

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
