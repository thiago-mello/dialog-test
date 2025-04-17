package deleteuser

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/database"
)

type DeleteUserUseCase interface {
	DeleteUser(ctx context.Context, command DeleteUserCommand) error
}

func NewDeleteUserUseCase(db *sqlx.DB, cache cache.Cache) DeleteUserUseCase {
	return &DeleteUserService{
		persistence: database.NewUsersDatabaseOutputPort(db),
		cache:       cache,
	}
}

type DeleteUserService struct {
	persistence database.UsersDatabaseOutputPort
	cache       cache.Cache
}

// DeleteUser deletes a user from the system by their ID and clears associated cache entries
// It performs the following steps:
// 1. Attempts to delete the user from the database using the provided user ID
// 2. Returns an error if the deletion fails or if no user was found
// 3. Clears any cached data for the deleted user using a pattern match
//
// Parameters:
//   - ctx: Context for the operation
//   - command: DeleteUserCommand containing the ID of the user to delete
//
// Returns:
//   - error: nil if successful, NotFoundError if user doesn't exist, or any database errors
func (d *DeleteUserService) DeleteUser(ctx context.Context, command DeleteUserCommand) error {
	rowsAffected, err := d.persistence.DeleteById(ctx, nil, command.UserId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errs.NotFoundError("user was not found")
	}

	// Evict cache for user
	userCacheKeysPattern := fmt.Sprintf("user:%s*", command.UserId.String())
	d.cache.DeleteByPattern(ctx, userCacheKeysPattern)

	return nil
}
