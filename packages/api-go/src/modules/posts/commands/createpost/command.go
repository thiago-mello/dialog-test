package createpost

import "github.com/google/uuid"

type CreatePostCommand struct {
	UserID   uuid.UUID
	Content  string
	IsPublic bool
}
