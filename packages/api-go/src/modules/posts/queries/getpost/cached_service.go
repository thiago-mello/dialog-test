package getpost

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"github.com/leandro-andrade-candido/api-go/src/libs/cache"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/database"
	"github.com/leandro-andrade-candido/api-go/src/modules/posts/domain"
)

type CachedGetPostService struct {
	baseService database.PostsDatabaseOutputPort
	cache       cache.Cache
	ttl         time.Duration
}

func NewCachedGetPostService(base database.PostsDatabaseOutputPort, cache cache.Cache) *CachedGetPostService {
	return &CachedGetPostService{
		baseService: base,
		cache:       cache,
		ttl:         1 * time.Hour,
	}
}

// FindByID retrieves a post by its ID, with caching support.
// It first attempts to fetch the post from cache using a composite key of user ID and post ID.
// If not found in cache, it queries the database and validates post ownership.
// Successfully retrieved posts are cached for future requests.
//
// Parameters:
//   - ctx: Context for the operation
//   - id: UUID of the post to retrieve
//   - userId: UUID of the user requesting the post
//
// Returns:
//   - *domain.Post: The retrieved post if found and owned by the user
//   - error: NotFoundError if post doesn't exist or belong to user, or other errors from cache/database
func (s *CachedGetPostService) FindByID(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*domain.Post, error) {
	cacheKey := fmt.Sprintf("user:%s;post:%s", userId.String(), id.String())

	var post *domain.Post
	if err := s.cache.Get(ctx, cacheKey, &post); err == nil {
		return post, nil
	}

	post, err := s.baseService.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// Checks if post belongs to current user
	if post == nil || post.UserID != userId {
		return nil, errs.NotFoundError("post not found")
	}

	if post != nil {
		s.cache.Set(ctx, cacheKey, post, s.ttl)
	}
	return post, nil
}
