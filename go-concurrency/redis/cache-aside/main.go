package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Simulated database.
var database = map[string]User{
	"1": {
		ID:   "1",
		Name: "Alice",
	},
}

func getUser(ctx context.Context, client *redis.Client, userID string) (User, error) {
	cacheKey := "user:" + userID

	// 1. Check Redis.
	cachedData, err := client.Get(ctx, cacheKey).Result()

	if err == nil {
		// Cache hit.
		var user User

		if err := json.Unmarshal([]byte(cachedData), &user); err != nil {
			return User{}, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}

		fmt.Println("Cache hit")
		return user, nil
	}

	if !errors.Is(err, redis.Nil) {
		// Unexpected Redis error.
		return User{}, fmt.Errorf("Redis GET failed: %w", err)
	}

	// 2. Cache miss: read from database.
	fmt.Println("Cache miss - reading from database")

	user, exists := database[userID]
	if !exists {
		return User{}, fmt.Errorf("user %s not found", userID)
	}

	// 3. Convert database result to JSON.
	userData, err := json.Marshal(user)
	if err != nil {
		return User{}, fmt.Errorf("failed to marshal user: %w", err)
	}

	// 4. Store result in Redis.
	if err := client.Set(ctx, cacheKey, userData, 5*time.Minute).Err(); err != nil {
		log.Println("warning: failed to populate cache:", err)
	}

	// 5. Return the user.
	return user, nil
}

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	defer client.Close()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	// First request.
	user, err := getUser(ctx, client, "1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	// Second request.
	user, err = getUser(ctx, client, "1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
