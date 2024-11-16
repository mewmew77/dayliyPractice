package pub_sub

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func PubMessage(ctx context.Context, client *redis.Client, channel string, msg string) {
	client.Publish(channel, msg)
}

func SubMessage(ctx context.Context, client *redis.Client, channels []string) {
	sub := client.Subscribe(channels...)
	_, err := sub.Receive()
	if err != nil {
		panic("Get Message failed.")
	}
	ch := sub.Channel()
	for msg := range ch {
		switch msg.Channel {
		case "ch1":
			fmt.Println("ch1: ", msg.Payload)
		case "ch2":
			fmt.Println("ch2: ", msg.Payload)
		default:
			fmt.Println("unknown channel ")
		}
	}
}
