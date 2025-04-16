package unlikepost

import "github.com/google/uuid"

type UnlikePostCommand struct {
	PostID uuid.UUID
	UserID uuid.UUID
}
