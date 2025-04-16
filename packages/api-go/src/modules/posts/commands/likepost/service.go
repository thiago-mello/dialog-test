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
