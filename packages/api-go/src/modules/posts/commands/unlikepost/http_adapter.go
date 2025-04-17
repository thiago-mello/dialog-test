package unlikepost

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
)

type UnlikePostHttpAdapter struct {
	useCase UnlikePostUseCase
}

func NewUnlikePostAdapter(db *sqlx.DB) *UnlikePostHttpAdapter {
	postsDB := database.NewPostsDatabaseOutputPort(db)
	likesDB := database.NewLikesDatabaseOutputPort(db)
	return &UnlikePostHttpAdapter{
		useCase: NewUseCase(postsDB, likesDB),
	}
}

// Handle UnlikePost godoc
// @Summary Removes Like from a post
// @Description Removes Like from a post by ID
// @Tags Posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 204
// @Failure 400 {object} errs.ErrorResponse
// @Failure 404 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /v1/posts/{id}/likes [delete]
// @Security ApiKeyAuth
func (a *UnlikePostHttpAdapter) Handle(c echo.Context) error {
	appCtx := c.(*context.ApplicationContext)
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errs.BadRequestError("invalid post ID")
	}

	command := UnlikePostCommand{
		PostID: postID,
		UserID: appCtx.User.Id,
	}

	if err := a.useCase.UnlikePost(c.Request().Context(), command); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
