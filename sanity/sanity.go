package sanity

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type sanityHandler struct {
	client *redis.Client
}

func sanityConstructor(redis *redis.Client) *sanityHandler {
	return &sanityHandler{
		client: redis,
	}
}

func (cfg *sanityHandler) sanity() {
	client := cfg.client

	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("NOT SANE", err)
	}

	fmt.Println("Redis:", pong)
}
