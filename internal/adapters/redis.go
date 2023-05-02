package adapters

import (
	"context"
	"fmt"
	"log"
	"users-app/domain"

	"github.com/go-redis/redis/v8"
)

// publisher is a struct used to manage connections to Redis and publish events
// to the specified Redis channel.
type publisher struct {
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

func NewPublisher(redisConfig RedisConfig) publisher {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}

	return publisher{client: client}
}

// PublishEvent publishes an event to the Redis channel. Currently, the event is serialized using
// the fmt.Sprintf function, but it is recommended to use proper event marshalling, such as protobuf or JSON.
func (p publisher) PublishEvent(ctx context.Context, event domain.Event) error {

	// todo - add proper event marshalling - protobuf, json, etc
	msg := fmt.Sprintf("%#v", event)

	return p.client.Publish(ctx, p.channel, msg).Err()
}
