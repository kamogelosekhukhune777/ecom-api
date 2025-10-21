package userbus

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/types/role"
)

type User struct {
	ID           uuid.UUID
	UserName     string
	FirstName    string
	LastName     string
	MobileNumber string
	Email        mail.Address
	PasswordHash []byte
	Enabled      bool
	Roles        []role.Role
	DateCreated  time.Time
	DateUpdated  time.Time
}

type NewUser struct {
	UserName     string
	FirstName    string
	LastName     string
	MobileNumber string
	Email        mail.Address
	Password     string
	Roles        []role.Role
}

type UpdateUser struct {
	UserName     *string
	FirstName    *string
	LastName     *string
	MobileNumber *string
	Email        *mail.Address
	Password     *string
	Roles        []role.Role
	Enabled      *bool
}
