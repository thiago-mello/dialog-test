package listposts

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/params"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/projections"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
	"github.com/samber/lo"
)

type ListPostsHttpAdapter struct {
	persistence database.PostsDatabaseOutputPort
}

func NewListPostAdapter(db *sqlx.DB) *ListPostsHttpAdapter {
	return &ListPostsHttpAdapter{
		persistence: database.NewPostsDatabaseOutputPort(db),
	}
}

func (a *ListPostsHttpAdapter) Query(ctx echo.Context) error {
	appCtx := ctx.(*context.ApplicationContext)

	queryParams := dto.ListPostRequestDto{}
	err := utils.BindAndValidate(ctx, &queryParams)
	if err != nil {
		return err
	}

	filters := params.GetPostsParams{
		PageSize:      utils.CalculatePageSize(int32(queryParams.PageSize)),
		LastSeenId:    utils.StringPointerToUuid(queryParams.LastSeenId),
		CurrentUserId: appCtx.User.Id,
	}

	posts, err := a.persistence.ListPosts(ctx.Request().Context(), filters)
	if err != nil {
		return err
	}

	mappedPosts := lo.Map(posts, func(post *projections.ListPostsProjection, _ int) dto.ListPostResponseDto {
		return dto.ListPostResponseDto{
			Id:      post.Id.String(),
			Content: post.Content,
			User: dto.ListPostUserDto{
				Id:   post.UserId.String(),
				Name: post.UserName,
				Bio:  post.UserBio,
			},
			CreatedAt:         post.CreatedAt,
			UpdatedAt:         post.UpdatedAt,
			LikeCount:         post.LikeCount,
			UserLikedThisPost: post.UserLikedThisPost,
		}
	})

	return ctx.JSON(http.StatusOK, mappedPosts)
}
