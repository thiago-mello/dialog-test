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
	persistence database.PostsDatabaseOutputPort
}

func NewUseCase(db *sqlx.DB) DeletePostUseCase {
	return &DeletePostService{
		persistence: database.NewPostsDatabaseOutputPort(db),
	}
}

func (s *DeletePostService) DeletePost(ctx context.Context, command DeletePostCommand) error {
	err := s.persistence.Delete(ctx, nil, command.PostID, command.UserID)

	if err != nil && err.Error() == "no posts deleted" {
		return errs.NotFoundError("Post not found")
	}

	return err
}
