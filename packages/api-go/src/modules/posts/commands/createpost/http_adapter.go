package createpost

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type CreatePostHttpAdapter struct {
	useCase CreatePostUseCase
	cache   cache.Cache
}

func NewCreatePostAdapter(db *sqlx.DB, cache cache.Cache) *CreatePostHttpAdapter {
	return &CreatePostHttpAdapter{useCase: NewUseCase(db), cache: cache}
}

// Handle CreatePost godoc
// @Summary Creates a new post
// @Description Creates a new post for the authenticated user
// @Tags Posts
// @Accept json
// @Produce json
// @Param payload body dto.CreatePostDto true "Post content"
// @Success 201 {object} dto.PostCreatedResponseDto
// @Failure 400 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /v1/posts [post]
// @Security ApiKeyAuth
func (c *CreatePostHttpAdapter) Handle(ctx echo.Context) error {
	appCtx := ctx.(*context.ApplicationContext)
	body := &dto.CreatePostDto{}

	if err := utils.BindAndValidate(ctx, body); err != nil {
		return err
	}

	sanitizedContent := utils.SanitizeHTML(body.Content)

	command := CreatePostCommand{
		UserID:   appCtx.User.Id,
		Content:  sanitizedContent,
		IsPublic: body.IsPublic,
	}

	post, err := c.useCase.CreatePost(ctx.Request().Context(), command)
	if err != nil {
		return err
	}

	c.invalidateCache(ctx, appCtx.User.Id, post.ID)
	return ctx.JSON(http.StatusCreated, dto.PostCreatedResponseDto{
		ID:       post.ID.String(),
		Content:  post.Content,
		IsPublic: post.IsPublic,
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
func (c *CreatePostHttpAdapter) invalidateCache(ctx echo.Context, userID uuid.UUID, postId uuid.UUID) {
	postKey := fmt.Sprintf("user:%s;post:%s", userID.String(), postId.String())
	timelinePattern := fmt.Sprintf("user:%s;timeline*", userID.String())
	c.cache.Delete(ctx.Request().Context(), postKey)
	c.cache.DeleteByPattern(ctx.Request().Context(), timelinePattern)
}
