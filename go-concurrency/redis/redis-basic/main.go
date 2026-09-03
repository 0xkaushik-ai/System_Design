package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9" // we used this becuz it's the latest version and in mod we did go get github.com/go-redis/redis/v8 there it get downloaded  and here it get imported  and we can use it in our code
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	// Is this a function or what
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// The meaning of the defer statement is that it will execute the function call (in this case, client.Close()) when the surrounding function (main) returns. This is useful for cleaning up resources, such as closing a database connection or a file, when they are no longer needed. In this case, it ensures that the Redis client connection is properly closed when the main function exits, preventing resource leaks and ensuring that the connection is released back to the system.
	defer client.Close()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	err := client.Set(ctx, "user:1:name", "kumail", time.Minute).Err()
	if err != nil {
		log.Fatal("Failed to set value in Redis:", err)
	}

	name, err := client.Get(ctx, "user:1:name").Result()
	if err != nil {
		log.Fatal("Failed to get value from Redis:", err)
	}
	fmt.Println("User name: ", name)
}
