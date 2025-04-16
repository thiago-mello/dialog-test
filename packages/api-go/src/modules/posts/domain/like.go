package domain

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	ID        uuid.UUID `db:"id"`
	PostID    uuid.UUID `db:"post_id"`
	UserID    uuid.UUID `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
