package listposts

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

type ListPostsHttpAdapter struct {
	cachedService *CachedListPostService
}

func NewListPostAdapter(db *sqlx.DB, cache cache.Cache) *ListPostsHttpAdapter {
	return &ListPostsHttpAdapter{
		cachedService: NewCachedListPostService(database.NewPostsDatabaseOutputPort(db), cache),
	}
}

// Query ListPosts godoc
// @Summary List posts
// @Description Lists posts from other users (public only)
// @Tags Posts
// @Produce json
// @Param page_size query int false "Page size (1-50)"
// @Param last_seen_id query string false "Last seen post ID"
// @Success 200 {array} dto.ListPostResponseDto
// @Failure 400 {object} errs.ErrorResponse
// @Failure 500 {object} errs.ErrorResponse
// @Router /v1/posts [get]
// @Security ApiKeyAuth
func (a *ListPostsHttpAdapter) Query(ctx echo.Context) error {
	appCtx := ctx.(*context.ApplicationContext)

	queryParams := dto.ListPostRequestDto{}
	err := utils.BindAndValidate(ctx, &queryParams)
	if err != nil {
		return err
	}

	posts, err := a.cachedService.ListPosts(ctx.Request().Context(), queryParams, appCtx.User.Id)

	return ctx.JSON(http.StatusOK, posts)
}
