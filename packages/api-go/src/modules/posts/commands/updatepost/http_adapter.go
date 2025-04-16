package updatepost

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type UpdatePostHttpAdapter struct {
	useCase UpdatePostUseCase
}

func NewUpdatePostAdapter(db *sqlx.DB) *UpdatePostHttpAdapter {
	return &UpdatePostHttpAdapter{useCase: NewUseCase(db)}
}

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

	return c.JSON(http.StatusOK, dto.PostUpdatedResponseDto{
		ID:       updatedPost.ID.String(),
		Content:  updatedPost.Content,
		IsPublic: updatedPost.IsPublic,
	})
}
