package main

import (
	"learn-redis/internal/limiter"
	"learn-redis/sanity"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	client *redis.Client
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	redisURL := os.Getenv("redis_url")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal(err)
	}

	clt := redis.NewClient(opt)

	var cfg Config

	cfg.client = clt

	sanityconfig := sanity.SanityConstructor(cfg.client)

	ptr := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: ptr,
	}

	ptr.HandleFunc("GET /health", sanityconfig.Sanity)
	ptr.HandleFunc("GET /checklimit/{userID}", limiter.CheckLimit)

	srv.ListenAndServe()

}
