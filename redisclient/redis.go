package redisclient

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client   *redis.Client
	initOnce sync.Once
)

func GetRedisClient() *redis.Client {
	initOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			log.Fatalf("Redis connection failed: %v", err)
		}
		log.Println("Connected to Redis")
	})

	return client
}
