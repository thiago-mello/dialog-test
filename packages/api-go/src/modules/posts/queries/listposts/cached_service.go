package listposts

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/libs/utils"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/params"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database/projections"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/dto"
	"github.com/samber/lo"
)

type CachedListPostService struct {
	baseService database.PostsDatabaseOutputPort
	cache       cache.Cache
	ttl         time.Duration
}

func NewCachedListPostService(base database.PostsDatabaseOutputPort, cache cache.Cache) *CachedListPostService {
	return &CachedListPostService{
		baseService: base,
		cache:       cache,
		ttl:         30 * time.Minute,
	}
}

func (s *CachedListPostService) ListPosts(ctx context.Context, filters dto.ListPostRequestDto, userId uuid.UUID) (*[]dto.ListPostResponseDto, error) {
	cacheKey := fmt.Sprintf("user:%s;timeline", userId.String())
	if filters.LastSeenId != nil {
		cacheKey += fmt.Sprintf(";last:%s", *filters.LastSeenId)
	}

	var posts *[]dto.ListPostResponseDto
	if err := s.cache.Get(ctx, cacheKey, &posts); err == nil {
		return posts, nil
	}

	dbParams := params.GetPostsParams{
		PageSize:      utils.CalculatePageSize(int32(filters.PageSize)),
		LastSeenId:    utils.StringPointerToUuid(filters.LastSeenId),
		CurrentUserId: userId,
		ShowPrivate:   false,
	}

	dbPosts, err := s.baseService.ListPosts(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	mappedPosts := lo.Map(dbPosts, func(post *projections.ListPostsProjection, _ int) dto.ListPostResponseDto {
		return dto.ListPostResponseDto{
			Id:        post.Id.String(),
			Content:   post.Content,
			IsPrivate: post.IsPrivate,
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

	// only cache if it has posts or if it is the last page of timeline
	if len(mappedPosts) > 0 || filters.LastSeenId != nil {
		s.cache.Set(ctx, cacheKey, mappedPosts, s.ttl)
	}
	return &mappedPosts, nil
}
