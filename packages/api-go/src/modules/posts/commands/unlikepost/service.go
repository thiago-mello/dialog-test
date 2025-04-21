package unlikepost

import (
	"context"

	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
)

type UnlikePostUseCase interface {
	UnlikePost(ctx context.Context, command UnlikePostCommand) error
}

type UnlikePostService struct {
	postsPersistence database.PostsDatabaseOutputPort
	likesPersistence database.LikesDatabaseOutputPort
}

func NewUseCase(postsDB database.PostsDatabaseOutputPort, likesDB database.LikesDatabaseOutputPort) UnlikePostUseCase {
	return &UnlikePostService{
		postsPersistence: postsDB,
		likesPersistence: likesDB,
	}
}

// UnlikePost removes a like from a post for a specific user
// ctx: The context for the operation
// command: Contains the PostID and UserID for the unlike operation
// Returns an error if the operation fails, nil on success
func (s *UnlikePostService) UnlikePost(ctx context.Context, command UnlikePostCommand) error {
	return s.likesPersistence.UnlikePost(ctx, command.PostID, command.UserID)
}
