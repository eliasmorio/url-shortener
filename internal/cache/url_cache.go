package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type UrlCache interface {
	Store(shortUrl string, longUrl string) error
	GetLongUrl(shortUrl string) (string, error)
}

type urlRedisCache struct {
	client *redis.Client
}

func NewUrlCache(redisClient *redis.Client) UrlCache {
	return &urlRedisCache{
		client: redisClient,
	}
}

func NewUrlCacheFromConfig(config RedisConfig) UrlCache {
	client := NewRedisClient(config)
	return NewUrlCache(client)
}

func (c *urlRedisCache) Store(shortUrl string, longUrl string) error {
	return c.client.Set(context.TODO(), shortUrl, longUrl, 0).Err()
}

func (c *urlRedisCache) GetLongUrl(shortUrl string) (string, error) {
	return c.client.Get(context.TODO(), shortUrl).Result()
}
