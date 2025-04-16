package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/sql"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type LikesDatabaseOutputAdapter struct {
	db *sqlx.DB
}

func NewLikesDatabaseOutputPort(db *sqlx.DB) LikesDatabaseOutputPort {
	return &LikesDatabaseOutputAdapter{db: db}
}

type LikesDatabaseOutputPort interface {
	LikePost(ctx context.Context, like *domain.PostLike) error
	UnlikePost(ctx context.Context, postID, userID uuid.UUID) error
}

func (l *LikesDatabaseOutputAdapter) LikePost(ctx context.Context, like *domain.PostLike) error {
	sqlString, err := sql.GetSql("post.Like", like)
	if err != nil {
		return err
	}

	_, err = l.db.NamedExecContext(ctx, sqlString, like)
	if err != nil {
		return err
	}

	return nil
}

func (l *LikesDatabaseOutputAdapter) UnlikePost(ctx context.Context, postID, userID uuid.UUID) error {
	sqlString, err := sql.GetSql("post.Unlike", nil)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"post_id": postID,
		"user_id": userID,
	}

	_, err = l.db.NamedExecContext(ctx, sqlString, params)
	return err
}
