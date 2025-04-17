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
