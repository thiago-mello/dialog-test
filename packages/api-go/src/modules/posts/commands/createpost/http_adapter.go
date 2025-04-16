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

// Handle processes HTTP requests to create a new post
// It performs the following steps:
// 1. Extracts the application context and validates request body
// 2. Sanitizes the post content to remove unsafe HTML
// 3. Creates a command with user ID, sanitized content and visibility setting
// 4. Calls the use case to create the post
// 5. Returns the created post details in JSON format
//
// Parameters:
//   - ctx: Echo context containing the HTTP request/response
//
// Returns:
//   - error: Any error that occurred during processing
//   - On success: Returns HTTP 201 with post details
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
