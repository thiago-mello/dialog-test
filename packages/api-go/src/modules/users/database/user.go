package database

import (
	"context"
	"errors"

	goSql "database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/sql"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/users/domain"
)

func NewUsersDatabaseOutputPort(db *sqlx.DB) UsersDatabaseOutputPort {
	return &UsersDatabaseOutputAdapter{db: db}
}

type UsersDatabaseOutputPort interface {
	Insert(ctx context.Context, tx *sqlx.Tx, user *domain.User) (bool, error)
	UpdateById(ctx context.Context, tx *sqlx.Tx, user *domain.User) (bool, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	DeleteById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (int64, error)
}

type UsersDatabaseOutputAdapter struct {
	db *sqlx.DB
}

func (u *UsersDatabaseOutputAdapter) Insert(ctx context.Context, tx *sqlx.Tx, user *domain.User) (userExists bool, err error) {
	dbTx := tx
	if tx == nil {
		dbTx = u.db.MustBeginTx(ctx, &goSql.TxOptions{})
	}

	sqlString, err := sql.GetSql("user.Insert", user)
	if err != nil {
		dbTx.Rollback()
		return false, err
	}

	_, err = dbTx.NamedExecContext(ctx, sqlString, user)
	if err != nil {
		dbTx.Rollback()
		return utils.IsConstraintViolation(err, "users_email_key"), err
	}

	// commits transaction if it was created inside this method
	if tx == nil {
		return false, dbTx.Commit()
	}

	return false, nil
}

func (u *UsersDatabaseOutputAdapter) UpdateById(ctx context.Context, tx *sqlx.Tx, user *domain.User) (userExists bool, err error) {
	dbTx := tx
	if tx == nil {
		dbTx = u.db.MustBeginTx(ctx, &goSql.TxOptions{})
	}

	sqlString, err := sql.GetSql("user.UpdateById", user)
	if err != nil {
		dbTx.Rollback()
		return false, err
	}

	result, err := dbTx.NamedExecContext(ctx, sqlString, user)
	if err != nil {
		dbTx.Rollback()
		return utils.IsConstraintViolation(err, "users_email_key"), err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		dbTx.Rollback()
		return false, errors.New("user not found")
	}

	// commits transaction if it was created inside this method
	if tx == nil {
		return false, dbTx.Commit()
	}

	return false, nil
}

func (u *UsersDatabaseOutputAdapter) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	sqlString, err := sql.GetSql("user.FindByEmail", nil)
	if err != nil {
		return nil, err
	}

	user := domain.User{}

	sqlString, args, err := utils.TranslateNamedQuery(sqlString, map[string]interface{}{"email": email})

	err = u.db.GetContext(ctx, &user, sqlString, args...)
	if err == goSql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteById deletes a user record from the database by their UUID
// Parameters:
//   - ctx: Context for the database operation
//   - tx: Optional transaction. If nil, a new transaction will be created
//   - id: UUID of the user to delete
//
// Returns:
//   - int64: Number of rows affected by the delete operation
//   - error: Any error that occurred during the operation
func (u *UsersDatabaseOutputAdapter) DeleteById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (int64, error) {
	dbTx := tx
	if tx == nil {
		dbTx = u.db.MustBeginTx(ctx, &goSql.TxOptions{})
	}

	sqlString, err := sql.GetSql("user.DeleteById", nil)
	if err != nil {
		dbTx.Rollback()
		return 0, err
	}

	result, err := dbTx.NamedExecContext(ctx, sqlString, map[string]interface{}{"id": id})
	if err != nil {
		dbTx.Rollback()
		return 0, err
	}

	// commits transaction if it was created inside this method
	if tx == nil {
		err = dbTx.Commit()
		if err != nil {
			return 0, err
		}
	}

	return result.RowsAffected()
}
