package database

import (
	"context"

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
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type UsersDatabaseOutputAdapter struct {
	db *sqlx.DB
}

func (u *UsersDatabaseOutputAdapter) Insert(ctx context.Context, tx *sqlx.Tx, user *domain.User) (userExists bool, err error) {
	dbTx := tx
	if tx == nil {
		dbTx = u.db.MustBegin()
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

func (u *UsersDatabaseOutputAdapter) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	sqlString, err := sql.GetSql("user.ExistsByEmail", nil)
	if err != nil {
		return false, err
	}

	rows, err := u.db.NamedQueryContext(ctx, sqlString, map[string]interface{}{"email": email})
	if rows.Err() != nil {
		return false, nil
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
