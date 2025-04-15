package updateuser

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
)

type UpdateUserUseCase interface {
	// Creates and persists new user into the database
	UpdateUser(ctx context.Context, command UpdateUserCommand) error
}

type UpdateUserService struct {
	persistence database.UsersDatabaseOutputPort
}

func NewUpdateUserUseCase(db *sqlx.DB) UpdateUserUseCase {
	return &UpdateUserService{
		persistence: database.NewUsersDatabaseOutputPort(db),
	}
}

// UpdateUser updates an existing user in the database based on the provided command
// It takes a context and UpdateUserCommand containing the user details to update
// If a password is provided, it will be hashed before storing
// Returns error if:
// - Password hashing fails
// - The requested email is already taken by another user
// - Database operation fails
func (u *UpdateUserService) UpdateUser(ctx context.Context, command UpdateUserCommand) error {
	user := &domain.User{
		ID:    command.UserId,
		Email: command.Email,
		Name:  command.Name,
		Bio:   command.Bio,
	}

	// only update password if user sends it
	if command.Password != "" {
		hash, err := utils.HashPassword(command.Password)
		if err != nil {
			return err
		}
		user.PasswordHash = hash
	}

	userExists, err := u.persistence.UpdateById(ctx, nil, user)
	if userExists {
		return errs.BadRequestError("The email requested is not available")
	}

	return err
}
