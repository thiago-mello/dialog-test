package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/sql"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type PostsDatabaseOutputAdapter struct {
	db *sqlx.DB
}

func NewPostsDatabaseOutputPort(db *sqlx.DB) PostsDatabaseOutputPort {
	return &PostsDatabaseOutputAdapter{db: db}
}

type PostsDatabaseOutputPort interface {
	Save(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error
}

// Save persists a post entity to the database
// Parameters:
//   - ctx: Context for the database operation
//   - tx: Optional transaction. If nil, a new transaction will be created
//   - post: The post domain entity to be saved
//
// Returns:
//   - error if the operation fails, nil on success
func (p *PostsDatabaseOutputAdapter) Save(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	sqlString, err := sql.GetSql("post.Save", post)
	if err != nil {
		return err
	}

	dbTx := tx
	if tx == nil {
		dbTx = p.db.MustBegin()
	}

	_, err = dbTx.NamedExecContext(ctx, sqlString, post)
	if err != nil {
		dbTx.Rollback()
		return err
	}

	if tx == nil {
		return dbTx.Commit()
	}
	return nil
}
