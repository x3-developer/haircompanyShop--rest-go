package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
}

type redisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService(ctx context.Context, addr, password string, db int) RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisService{
		client: rdb,
		ctx:    ctx,
	}
}

func (r *redisService) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *redisService) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", fmt.Errorf("key does not exist")
	}
	return val, err
}

func (r *redisService) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *redisService) Exists(key string) (bool, error) {
	count, err := r.client.Exists(r.ctx, key).Result()
	return count > 0, err
}
