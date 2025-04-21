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
	Persistence database.UsersDatabaseOutputPort
}

func NewUpdateUserUseCase(db *sqlx.DB) UpdateUserUseCase {
	return &UpdateUserService{
		Persistence: database.NewUsersDatabaseOutputPort(db),
	}
}

// UpdateUser updates an existing user's information in the database
// It takes a context and UpdateUserCommand containing the user details to update
// Returns error if:
// - Password hashing fails
// - Email is already taken by another user
// - User is not found
// - Any other database error occurs
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

	userExists, err := u.Persistence.UpdateById(ctx, nil, user)
	if userExists {
		return errs.BadRequestError("The email requested is not available")
	}

	if err != nil && err.Error() == "user not found" {
		return errs.NotFoundError("User not found")
	}

	return err
}
