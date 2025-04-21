package listmyposts

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

type CachedListMyPostsService struct {
	baseService database.PostsDatabaseOutputPort
	cache       cache.Cache
	ttl         time.Duration
}

func NewCachedListMyPostsService(base database.PostsDatabaseOutputPort, cache cache.Cache) *CachedListMyPostsService {
	return &CachedListMyPostsService{
		baseService: base,
		cache:       cache,
		ttl:         30 * time.Minute,
	}
}

// ListPosts retrieves a list of posts for a specific user with caching support
//
// Parameters:
//   - ctx: The context for the request
//   - filters: Filtering options including page size and last seen post ID
//   - userId: UUID of the user whose posts are being retrieved
//
// Returns:
//   - *[]dto.ListPostResponseDto: Pointer to slice of post response DTOs
//   - error: Error if retrieval fails
//
// The function first attempts to fetch posts from cache using a key based on userId
// and lastSeenId. If not found in cache, it queries the database with the provided
// parameters, maps the database results to DTOs, and stores them in cache before
// returning.
func (s *CachedListMyPostsService) ListPosts(ctx context.Context, filters dto.ListPostRequestDto, userId uuid.UUID) (*[]dto.ListPostResponseDto, error) {
	cacheKey := fmt.Sprintf("user:%s;timeline-my-posts", userId.String())
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
		UserId:        &userId,
		ShowPrivate:   true,
	}

	dbPosts, err := s.baseService.ListPosts(ctx, dbParams)
	if err != nil {
		return nil, err
	}

	mappedPosts := lo.Map(*dbPosts, func(post projections.ListPostsProjection, _ int) dto.ListPostResponseDto {
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

	if len(mappedPosts) > 0 || filters.LastSeenId != nil {
		s.cache.Set(ctx, cacheKey, mappedPosts, s.ttl)
	}
	return &mappedPosts, nil
}
