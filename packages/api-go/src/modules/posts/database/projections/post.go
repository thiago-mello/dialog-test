package projections

import (
	"time"

	"github.com/google/uuid"
)

type ListPostsProjection struct {
	Id                uuid.UUID `db:"id"`
	Content           string    `db:"content"`
	IsPrivate         bool      `db:"is_private"`
	UserId            uuid.UUID `db:"user_id"`
	UserName          string    `db:"user_name"`
	UserBio           *string   `db:"user_bio"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
	LikeCount         int32     `db:"likes"`
	UserLikedThisPost bool      `db:"user_liked"`
}
