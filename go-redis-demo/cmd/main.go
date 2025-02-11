package main

import (
	"context"
	"fmt"
	"go-redis-demo/pkg/db"
	"log"
)

func main() {
	fmt.Println("Program started....")

	redisClient, err := db.NewRedisClient()
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	defer redisClient.Client.Close() // Ensure proper cleanup on shutdow

	ctx := context.Background()
	err = redisClient.Client.Set(ctx, "message", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("Failed to set key in Redis: %v", err)
	}

	// Retrieve the key
	val, err := redisClient.Client.Get(ctx, "message").Result()
	if err != nil {
		log.Fatalf("Failed to get key from Redis: %v", err)
	}

	fmt.Println("Retrieved from Redis:", val)
}
