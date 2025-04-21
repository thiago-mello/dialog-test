package likepost

import (
	"context"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type LikePostUseCase interface {
	LikePost(ctx context.Context, command LikePostCommand) error
}

type LikePostService struct {
	postsPersistence database.PostsDatabaseOutputPort
	likesPersistence database.LikesDatabaseOutputPort
}

func NewUseCase(postsDB database.PostsDatabaseOutputPort, likesDB database.LikesDatabaseOutputPort) LikePostUseCase {
	return &LikePostService{
		postsPersistence: postsDB,
		likesPersistence: likesDB,
	}
}

// LikePost handles the process of liking a post
// It performs the following steps:
// 1. Verifies if the post exists in the database
// 2. Generates a new UUID for the like
// 3. Creates a new PostLike entity
// 4. Persists the like in the database
// Parameters:
//   - ctx: Context for the operation
//   - command: Contains PostID and UserID for the like operation
//
// Returns:
//   - error: Returns nil if successful, NotFoundError if post doesn't exist,
//     or any other error that occurred during the operation
func (s *LikePostService) LikePost(ctx context.Context, command LikePostCommand) error {
	post, err := s.postsPersistence.FindByID(ctx, command.PostID)
	if err != nil {
		return err
	}
	if post == nil {
		return errs.NotFoundError("post not found")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	like := &domain.PostLike{
		ID:     id,
		PostID: command.PostID,
		UserID: command.UserID,
	}

	return s.likesPersistence.LikePost(ctx, like)
}
