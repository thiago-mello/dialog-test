package updatepost

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type UpdatePostUseCase interface {
	UpdatePost(ctx context.Context, command UpdatePostCommand) (*domain.Post, error)
}

type UpdatePostService struct {
	persistence database.PostsDatabaseOutputPort
}

func NewUseCase(db *sqlx.DB) UpdatePostUseCase {
	return &UpdatePostService{
		persistence: database.NewPostsDatabaseOutputPort(db),
	}
}

// UpdatePost updates an existing post in the system
// It performs the following steps:
// 1. Finds the post by ID
// 2. Verifies the post exists and belongs to the requesting user
// 3. Updates the post content and visibility
// 4. Validates the updated post
// 5. Persists the changes
// Returns the updated post or an error if any step fails
func (s *UpdatePostService) UpdatePost(ctx context.Context, command UpdatePostCommand) (*domain.Post, error) {
	existingPost, err := s.persistence.FindByID(ctx, command.PostID)
	if err != nil {
		return nil, err
	}
	if existingPost == nil {
		return nil, errs.NotFoundError("post not found")
	}
	if existingPost.UserID != command.UserID {
		return nil, errs.NotFoundError("post not found")
	}

	existingPost.Content = command.Content
	existingPost.IsPublic = command.IsPublic

	if err := existingPost.Validate(); err != nil {
		return nil, errs.BadRequestError(err.Error())
	}

	if err := s.persistence.Update(ctx, nil, existingPost); err != nil {
		return nil, err
	}

	return existingPost, nil
}
