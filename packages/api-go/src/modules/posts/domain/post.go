package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/ddd"
)

type Post struct {
	ID       uuid.UUID
	UserID   uuid.UUID `db:"user_id"`
	Content  string
	IsPublic bool `db:"is_public"`
	ddd.AuditableModel
}

// Validate checks if the Post content meets the required criteria:
// - Content must not be empty
// - Content must not exceed 12000 characters
// Returns an error if validation fails, nil otherwise
func (p *Post) Validate() error {
	if len(p.Content) == 0 {
		return errors.New("content cannot be empty")
	}
	if len(p.Content) > 12000 {
		return errors.New("content exceeds 12000 characters")
	}
	return nil
}
