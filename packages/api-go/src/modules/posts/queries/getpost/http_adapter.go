package getpost

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type GetPostHttpAdapter struct {
	cachedService *CachedGetPostService
}

func NewGetPostAdapter(db *sqlx.DB, cache cache.Cache) *GetPostHttpAdapter {
	return &GetPostHttpAdapter{
		cachedService: NewCachedGetPostService(database.NewPostsDatabaseOutputPort(db), cache),
	}
}

func (a *GetPostHttpAdapter) Query(c echo.Context) error {
	appCtx := c.(*context.ApplicationContext)
	postID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errs.BadRequestError("invalid post ID")
	}

	post, err := a.cachedService.FindByID(c.Request().Context(), postID, appCtx.User.Id)
	if err != nil {
		return err
	}
	if post == nil {
		return errs.NotFoundError("post not found")
	}

	return c.JSON(http.StatusOK, dto.PostResponseDto{
		ID:        post.ID.String(),
		Content:   post.Content,
		IsPublic:  post.IsPublic,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}
