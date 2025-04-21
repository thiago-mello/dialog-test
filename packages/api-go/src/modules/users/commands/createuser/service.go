package createuser

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
)

type CreateUserUseCase interface {
	// Creates and persists new user into the database
	CreateNewUser(ctx context.Context, command CreateUserCommand) error
}

type CreateUserService struct {
	Persistence database.UsersDatabaseOutputPort
}

func NewUseCase(db *sqlx.DB) CreateUserUseCase {
	persistence := database.NewUsersDatabaseOutputPort(db)

	return &CreateUserService{Persistence: persistence}
}

// CreateNewUser creates a new user in the system. It performs the following steps:
// 1. Maps the command data to a domain user object
// 2. Validates the user data
// 3. Attempts to insert the user into the database
//
// Parameters:
//   - ctx: Context for the operation
//   - command: CreateUserCommand containing the user data
//
// Returns:
//   - error: BadRequestError if validation fails or email already exists
//   - error: Any other errors that occur during processing
func (c *CreateUserService) CreateNewUser(ctx context.Context, command CreateUserCommand) error {
	user, err := mapCommandToDomain(command)
	if err != nil {
		return err
	}

	if err := user.Validate(); err != nil {
		return errs.BadRequestError(err.Error())
	}

	userExists, err := c.Persistence.Insert(ctx, nil, user)
	if userExists {
		return errs.BadRequestError("The email requested is not available")
	}

	return err
}

// mapCommandToDomain converts a CreateUserCommand into a domain.User object
// It generates a new UUID for the user and hashes their password
//
// Parameters:
//   - command: CreateUserCommand containing the raw user data (email, name, bio, password)
//
// Returns:
//   - *domain.User: A new domain User object with generated ID and hashed password
//   - error: Error if UUID generation or password hashing fails
func mapCommandToDomain(command CreateUserCommand) (*domain.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	passwordHash, err := utils.HashPassword(command.Password)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           id,
		Email:        command.Email,
		Name:         command.Name,
		Bio:          command.Bio,
		PasswordHash: passwordHash,
	}, err
}
