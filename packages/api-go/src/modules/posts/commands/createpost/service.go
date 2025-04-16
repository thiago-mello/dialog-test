package createpost

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type CreatePostUseCase interface {
	CreatePost(ctx context.Context, command CreatePostCommand) (*domain.Post, error)
}

type CreatePostService struct {
	persistence database.PostsDatabaseOutputPort
}

func NewUseCase(db *sqlx.DB) CreatePostUseCase {
	return &CreatePostService{
		persistence: database.NewPostsDatabaseOutputPort(db),
	}
}

// CreatePost creates a new post in the system
// Parameters:
//   - ctx: Context for the operation
//   - command: CreatePostCommand containing the post details (UserID, Content, IsPublic)
//
// Returns:
//   - *domain.Post: The created post if successful
//   - error: BadRequestError if validation fails or any other error during save
func (c *CreatePostService) CreatePost(ctx context.Context, command CreatePostCommand) (*domain.Post, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	post := &domain.Post{
		ID:       id,
		UserID:   command.UserID,
		Content:  command.Content,
		IsPublic: command.IsPublic,
	}

	if err := post.Validate(); err != nil {
		return nil, errs.BadRequestError(err.Error())
	}

	if err := c.persistence.Save(ctx, nil, post); err != nil {
		return nil, err
	}

	return post, nil
}
