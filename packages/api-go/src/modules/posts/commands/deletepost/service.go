package deletepost

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
)

type DeletePostUseCase interface {
	DeletePost(ctx context.Context, command DeletePostCommand) error
}
type DeletePostService struct {
	Persistence database.PostsDatabaseOutputPort
}

func NewUseCase(db *sqlx.DB) DeletePostUseCase {
	return &DeletePostService{
		Persistence: database.NewPostsDatabaseOutputPort(db),
	}
}

// DeletePost deletes a post from the database based on the provided command
// It takes a context and DeletePostCommand containing PostID and UserID
// Returns NotFoundError if no post was found to delete
// Returns any other errors encountered during deletion
func (s *DeletePostService) DeletePost(ctx context.Context, command DeletePostCommand) error {
	err := s.Persistence.Delete(ctx, nil, command.PostID, command.UserID)

	if err != nil && err.Error() == "no posts deleted" {
		return errs.NotFoundError("Post not found")
	}

	return err
}
