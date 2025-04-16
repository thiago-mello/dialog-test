package likepost

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
)

type LikePostHttpAdapter struct {
	useCase LikePostUseCase
}

func NewLikePostAdapter(db *sqlx.DB) *LikePostHttpAdapter {
	postsDB := database.NewPostsDatabaseOutputPort(db)
	likesDB := database.NewLikesDatabaseOutputPort(db)
	return &LikePostHttpAdapter{
		useCase: NewUseCase(postsDB, likesDB),
	}
}

func (a *LikePostHttpAdapter) Handle(c echo.Context) error {
	appCtx := c.(*context.ApplicationContext)
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errs.BadRequestError("invalid post ID")
	}

	command := LikePostCommand{
		PostID: postID,
		UserID: appCtx.User.Id,
	}

	if err := a.useCase.LikePost(c.Request().Context(), command); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
