package adapters

import (
	"context"
	"fmt"
	"log"
	"users-app/domain"

	"github.com/go-redis/redis/v8"
)

type pubSub struct {
	client  *redis.Client
	channel string
}

type RedisConfig struct {
	Host          string
	Port          string
	Password      string
	DB            int
	EventsChannel string
}

func NewPubSub(redisConfig RedisConfig) pubSub {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to Redis")
	return pubSub{client: client}
}

// PublishEvent todo - refactor to domain Event interface
func (p pubSub) PublishEvent(ctx context.Context, event domain.Event) error {

	// todo - add proper event marshalling - protobuf, json, etc
	msg := fmt.Sprintf("%#v", event)
	return p.client.Publish(context.Background(), p.channel, msg).Err()
}
