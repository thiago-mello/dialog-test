package deletepost

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
)

type DeletePostHttpAdapter struct {
	useCase DeletePostUseCase
	cache   cache.Cache
}

func NewDeletePostAdapter(db *sqlx.DB, cache cache.Cache) *DeletePostHttpAdapter {
	return &DeletePostHttpAdapter{useCase: NewUseCase(db), cache: cache}
}

// Handle DeletePost godoc
// @Summary Deletes a post
// @Description Deletes a post by ID, only if it belongs to the authenticated user
// @Tags Posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 204
// @Failure 400 {object} errs.ErrorResponse
// @Failure 404 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /v1/posts/{id} [delete]
// @Security ApiKeyAuth
func (a *DeletePostHttpAdapter) Handle(c echo.Context) error {
	appCtx := c.(*context.ApplicationContext)
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errs.BadRequestError("invalid post ID")
	}
	command := DeletePostCommand{
		PostID: postID,
		UserID: appCtx.User.Id,
	}
	if err := a.useCase.DeletePost(c.Request().Context(), command); err != nil {
		return err
	}

	a.invalidateCache(c, appCtx.User.Id, postID)
	return c.NoContent(http.StatusNoContent)
}

func (a *DeletePostHttpAdapter) invalidateCache(ctx echo.Context, userID uuid.UUID, postId uuid.UUID) {
	postKey := fmt.Sprintf("user:%s;post:%s", userID.String(), postId.String())
	timelinePattern := fmt.Sprintf("user:%s;timeline*", userID.String())
	a.cache.Delete(ctx.Request().Context(), postKey)
	a.cache.DeleteByPattern(ctx.Request().Context(), timelinePattern)
}
