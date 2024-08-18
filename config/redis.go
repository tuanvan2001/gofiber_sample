package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
)

var ctx = context.Background()

var Redis *redis.Client

func ConnectRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	Redis = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	pong, err := Redis.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Printf("Redis connected: %v", pong)
}
