package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"tmdb-cli/internal/models"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr, password string, db int) (*RedisCache, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
		ctx:    ctx,
	}, nil
}

func (rc *RedisCache) GetMovies(key string) ([]models.Movie, error) {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	if err := json.Unmarshal([]byte(val), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (rc *RedisCache) SetMovies(key string, movies []models.Movie, expiration time.Duration) error {
	data, err := json.Marshal(movies)
	if err != nil {
		return err
	}
	return rc.client.Set(rc.ctx, key, data, expiration).Err()
}
