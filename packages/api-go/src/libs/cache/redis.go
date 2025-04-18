package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	err := redisotel.InstrumentTracing(client)
	if err != nil {
		log.Default().Println("failed to instrument redis tracing", err)
	}

	return &RedisCache{
		client: client,
	}
}

// Get retrieves a value from Redis for the given key and scans it into the dest interface.
// Returns an error if the operation fails or if the key doesn't exist.
func (r *RedisCache) Get(ctx context.Context, key string, dest any) error {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dest)
}

// Set stores a value in Redis with the specified key and time-to-live duration.
// Returns an error if the operation fails.
func (r *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

// Delete removes one or more keys from Redis.
// Returns an error if the operation fails.
func (r *RedisCache) Delete(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// DeleteByPattern removes all keys matching the specified pattern from Redis.
// It uses SCAN to iterate through keys in batches and deletes them using pipelining for better performance.
// Returns an error if scanning or deletion fails.
func (r *RedisCache) DeleteByPattern(ctx context.Context, pattern string) error {
	const batchSize = 100
	var cursor uint64
	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, pattern, batchSize).Result()
		if err != nil {
			return fmt.Errorf("error during SCAN: %w", err)
		}

		if len(keys) > 0 {
			// Use pipeline to delete keys in batch
			pipe := r.client.Pipeline()
			for _, key := range keys {
				pipe.Del(ctx, key)
			}
			_, err := pipe.Exec(ctx)
			if err != nil {
				return fmt.Errorf("error during pipeline execution: %w", err)
			}
		}

		// Move to the next cursor
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}

// Close terminates the connection to Redis.
// Returns an error if the operation fails.
func (r *RedisCache) Close() error {
	return r.client.Close()
}
