package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// RedisClient is the global Redis client.
// Other packages will use this to store and retrieve data from Redis.
var RedisClient *redis.Client

// ConnectRedis reads the Redis connection details from environment variables,
// creates a client, and checks that the connection is working.
func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port), // e.g. "localhost:6379"
	})

	// Ping Redis to make sure the connection is alive
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis: ", err)
	}

	log.Println("Connected to Redis successfully")
}
