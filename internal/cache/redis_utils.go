package cache

import (
	"UrlShortener/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisConfig struct {
	Url      string `env:"REDIS_URL" envDefault:"localhost:6379"`
	Username string `env:"REDIS_USER"`
	Password string `env:"REDIS_PASS"`
	Db       int    `env:"REDIS_DB" envDefault:"0"`
}

func GetRedisConfig() RedisConfig {
	cacheConfig := RedisConfig{}
	err := config.LoadConfig(&cacheConfig)
	if err != nil {
		panic(err)
	}
	return cacheConfig
}

func NewRedisClient(config RedisConfig) *redis.Client {
	slog.Info("Connecting to redis")
	client := redis.NewClient(&redis.Options{
		Addr:     config.Url,
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
	})
	err := client.Ping(context.TODO()).Err()
	if err != nil {
		slog.Error("Error while connecting to redis", err)
		panic(err)
	}
	slog.Info("Successfully connected to redis")
	return client

}
