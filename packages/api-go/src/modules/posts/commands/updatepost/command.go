package updatepost

import "github.com/google/uuid"

type UpdatePostCommand struct {
	PostID   uuid.UUID
	UserID   uuid.UUID
	Content  string
	IsPublic bool
}
