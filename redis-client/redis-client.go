package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

// redis-client is a small program that connects to Redis and subscribes to the events channel.
// It's used to test the Redis connection and for development of the app
func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server to check the connection
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	pubsub := client.Subscribe(context.Background(), os.Getenv("REDIS_EVENTS_CHANNEL"))
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received message: %s\n", msg.Payload)
	}
}
