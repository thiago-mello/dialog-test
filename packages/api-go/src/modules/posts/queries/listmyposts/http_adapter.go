package listmyposts

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
)

type ListMyPostsHttpAdapter struct {
	cachedService *CachedListMyPostsService
}

func NewListMyPostAdapter(db *sqlx.DB, cache cache.Cache) *ListMyPostsHttpAdapter {
	return &ListMyPostsHttpAdapter{
		cachedService: NewCachedListMyPostsService(database.NewPostsDatabaseOutputPort(db), cache),
	}
}

func (a *ListMyPostsHttpAdapter) Query(ctx echo.Context) error {
	appCtx := ctx.(*context.ApplicationContext)

	queryParams := dto.ListPostRequestDto{}
	err := utils.BindAndValidate(ctx, &queryParams)
	if err != nil {
		return err
	}

	posts, err := a.cachedService.ListPosts(ctx.Request().Context(), queryParams, appCtx.User.Id)

	return ctx.JSON(http.StatusOK, posts)
}
