package createpost

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type CreatePostHttpAdapter struct {
	useCase CreatePostUseCase
}

func NewCreatePostAdapter(db *sqlx.DB) *CreatePostHttpAdapter {
	return &CreatePostHttpAdapter{useCase: NewUseCase(db)}
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

	return ctx.JSON(http.StatusCreated, dto.PostCreatedResponseDto{
		ID:       post.ID.String(),
		Content:  post.Content,
		IsPublic: post.IsPublic,
	})
}
