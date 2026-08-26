package sanity

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type sanityHandler struct {
	client *redis.Client
}

func SanityConstructor(redis *redis.Client) *sanityHandler {
	return &sanityHandler{
		client: redis,
	}
}

func (cfg *sanityHandler) Sanity(w http.ResponseWriter, r *http.Request) {
	client := cfg.client

	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("NOT SANE", err)
	}

	fmt.Println("Redis:", pong)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
