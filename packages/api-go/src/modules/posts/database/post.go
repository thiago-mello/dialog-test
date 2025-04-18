package database

import (
	"context"
	sqlDb "database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/sql"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/params"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/projections"
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
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Post, error)
	Update(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error
	ListPosts(ctx context.Context, filters params.GetPostsParams) ([]*projections.ListPostsProjection, error)
	Delete(ctx context.Context, tx *sqlx.Tx, postID, userID uuid.UUID) error
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
		dbTx = p.db.MustBeginTx(ctx, &sqlDb.TxOptions{})
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

func (p *PostsDatabaseOutputAdapter) FindByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	sqlString, err := sql.GetSql("post.FindByID", nil)
	if err != nil {
		return nil, err
	}
	post := domain.Post{}
	query, args, err := utils.TranslateNamedQuery(sqlString, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	err = p.db.GetContext(ctx, &post, query, args...)
	if err == sqlDb.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *PostsDatabaseOutputAdapter) Update(ctx context.Context, tx *sqlx.Tx, post *domain.Post) error {
	sqlString, err := sql.GetSql("post.Update", post)
	if err != nil {
		return err
	}
	dbTx := tx
	if tx == nil {
		dbTx = p.db.MustBeginTx(ctx, &sqlDb.TxOptions{})
	}
	result, err := dbTx.NamedExecContext(ctx, sqlString, post)
	if err != nil {
		dbTx.Rollback()
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		dbTx.Rollback()
		return errors.New("no rows updated")
	}
	if tx == nil {
		return dbTx.Commit()
	}
	return nil
}

func (p *PostsDatabaseOutputAdapter) ListPosts(ctx context.Context, filters params.GetPostsParams) ([]*projections.ListPostsProjection, error) {
	sqlString, err := sql.GetSql("post.ListPosts", filters)
	if err != nil {
		return nil, err
	}
	posts := []*projections.ListPostsProjection{}
	query, args, err := utils.TranslateNamedQuery(sqlString, filters)
	if err != nil {
		return nil, err
	}
	err = p.db.SelectContext(ctx, &posts, query, args...)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostsDatabaseOutputAdapter) Delete(ctx context.Context, tx *sqlx.Tx, postID, userID uuid.UUID) error {
	sqlString, err := sql.GetSql("post.Delete", nil)
	if err != nil {
		return err
	}

	dbTx := tx
	if tx == nil {
		dbTx = p.db.MustBeginTx(ctx, &sqlDb.TxOptions{})
	}

	params := map[string]any{
		"id":      postID,
		"user_id": userID,
	}

	result, err := dbTx.NamedExecContext(ctx, sqlString, params)
	if err != nil {
		dbTx.Rollback()
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		dbTx.Rollback()
		return errors.New("no posts deleted")
	}

	if tx == nil {
		return dbTx.Commit()
	}
	return nil
}
