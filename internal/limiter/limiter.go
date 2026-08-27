package limiter

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
)

type limitHandler struct {
	client *redis.Client
}

//go:embed script.lua
var rateLimitScriptst string

var rateLimitScript = redis.NewScript(rateLimitScriptst)

func LimitConstructor(redis *redis.Client) *limitHandler {
	return &limitHandler{
		client: redis,
	}
}

func (cfg *limitHandler) CheckLimit(w http.ResponseWriter, r *http.Request) {

	userID := r.PathValue("userID")
	reqID, err := uuid.NewV7()
	if err != nil {
		log.Fatal("cant generate request iD: ", err)
	}

	ctx := r.Context()

	res, err := rateLimitScript.Run(ctx, cfg.client, []string{userID}, reqID.String()).Result()
	if err != nil {
		log.Fatal("something went wrong", err)
	}

	if res == 1 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(429)
	w.Write([]byte("hol up mate"))

}
