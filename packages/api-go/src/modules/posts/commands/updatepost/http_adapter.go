package updatepost

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type UpdatePostHttpAdapter struct {
	useCase UpdatePostUseCase
	cache   cache.Cache
}

func NewUpdatePostAdapter(db *sqlx.DB, cache cache.Cache) *UpdatePostHttpAdapter {
	return &UpdatePostHttpAdapter{useCase: NewUseCase(db), cache: cache}
}

// Handle UpdatePost godoc
// @Summary Updates a post
// @Description Updates a post's content and visibility
// @Tags Posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param payload body dto.UpdatePostDto true "Updated post content"
// @Success 200 {object} dto.PostUpdatedResponseDto
// @Failure 400 {object} errs.ErrorResponse
// @Failure 404 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /v1/posts/{id} [put]
// @Security ApiKeyAuth
func (a *UpdatePostHttpAdapter) Handle(c echo.Context) error {
	appCtx := c.(*context.ApplicationContext)
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errs.BadRequestError("invalid post ID")
	}

	body := &dto.UpdatePostDto{}
	if err := utils.BindAndValidate(c, body); err != nil {
		return err
	}

	sanitizedContent := utils.SanitizeHTML(body.Content)

	command := UpdatePostCommand{
		PostID:   postID,
		UserID:   appCtx.User.Id,
		Content:  sanitizedContent,
		IsPublic: body.IsPublic,
	}

	updatedPost, err := a.useCase.UpdatePost(c.Request().Context(), command)
	if err != nil {
		return err
	}

	// cache invalidation
	a.invalidateCache(c, appCtx.User.Id, postID)

	return c.JSON(http.StatusOK, dto.PostUpdatedResponseDto{
		ID:       updatedPost.ID.String(),
		Content:  updatedPost.Content,
		IsPublic: updatedPost.IsPublic,
	})
}

// invalidateCache removes the cache entries for a specific post and user timeline
// after a post is updated. This ensures that subsequent requests will fetch fresh data.
//
// Parameters:
//   - ctx: The Echo context containing the request information
//   - userID: UUID of the user who owns the post
//   - postId: UUID of the post being updated
//
// The function invalidates:
//   - The specific post cache entry
//   - All timeline entries for the user
func (a *UpdatePostHttpAdapter) invalidateCache(ctx echo.Context, userID uuid.UUID, postId uuid.UUID) {
	postKey := fmt.Sprintf("user:%s;post:%s", userID.String(), postId.String())
	timelinePattern := fmt.Sprintf("user:%s;timeline*", userID.String())
	a.cache.Delete(ctx.Request().Context(), postKey)
	a.cache.DeleteByPattern(ctx.Request().Context(), timelinePattern)
}
