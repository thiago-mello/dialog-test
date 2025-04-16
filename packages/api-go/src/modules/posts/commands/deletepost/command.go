package deletepost

import "github.com/google/uuid"

type DeletePostCommand struct {
	PostID uuid.UUID
	UserID uuid.UUID
}
