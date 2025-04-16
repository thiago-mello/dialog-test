package deletepost

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
)

type DeletePostHttpAdapter struct {
	useCase DeletePostUseCase
}

func NewDeletePostAdapter(db *sqlx.DB) *DeletePostHttpAdapter {
	return &DeletePostHttpAdapter{useCase: NewUseCase(db)}
}

// Handle processes a request to delete a post
// It extracts the post ID from the URL parameters and the user ID from the context
// Then calls the delete post use case and returns 204 No Content on success
// Returns 400 Bad Request if the post ID is invalid
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
	return c.NoContent(http.StatusNoContent)
}
