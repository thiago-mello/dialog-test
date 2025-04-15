package domain

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/ddd"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string `db:"password_hash"`
	Name         string
	Bio          *string
	ddd.AuditableModel
}

// Validate performs validation checks on the User struct fields
// It checks:
// - Email is not empty and is a valid email format
// - Password hash is not empty
// - Name is not empty
// Returns an error if any validation fails, nil otherwise
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is mandatory")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email")
	}

	if u.PasswordHash == "" {
		return errors.New("invalid password hash")
	}

	if u.Name == "" {
		return errors.New("name is mandatory")
	}

	return nil
}
