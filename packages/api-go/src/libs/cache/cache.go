package cache

import (
	"context"
	"time"
)

// Cache defines the interface for a generic caching system
type Cache interface {
	// Get retrieves a value from the cache by key and stores it in dest
	// Returns error if key not found or on retrieval failure
	Get(ctx context.Context, key string, dest any) error

	// Set stores a value in the cache with the given key and TTL duration
	// Returns error if the set operation fails
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	// Delete removes one or more keys from the cache
	// Returns error if the delete operation fails
	Delete(ctx context.Context, keys ...string) error

	// DeleteByPattern removes all keys matching the given pattern
	// Returns error if the delete operation fails
	DeleteByPattern(ctx context.Context, pattern string) error

	// Close cleans up cache resources and closes connections
	// Returns error if cleanup fails
	Close() error
}
