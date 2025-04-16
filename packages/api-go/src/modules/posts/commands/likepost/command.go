package likepost

import "github.com/google/uuid"

type LikePostCommand struct {
	PostID uuid.UUID
	UserID uuid.UUID
}
