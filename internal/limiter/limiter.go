package limiter

import (
	_ "embed"
	"log"
	"net/http"
	_ "strconv"
	_ "time"

	"github.com/redis/go-redis/v9"
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
	// reqID, err := uuid.NewV7()
	// if err != nil {
	// 	log.Fatal("cant generate request iD: ", err)
	// }

	ctx := r.Context()
	// seconds := time.Now().Unix()
	// window := seconds - 60
	// maxInclusive := strconv.Itoa(int(window))

	res, err := rateLimitScript.Run(ctx, cfg.client, []string{userID}).Result()
	if err != nil {
		log.Fatal("something went wrong", err)
	}

	// var res int
	// cfg.client.ZRemRangeByScore(ctx, userID, "-inf", maxInclusive)
	// if cfg.client.ZCard(ctx, userID).Val() >= 10 {
	// 	res = 0
	// } else {
	// 	res = 1
	// 	cfg.client.ZAdd(ctx, userID, redis.Z{
	// 		Score:  float64(seconds),
	// 		Member: reqID.String(),
	// 	})
	// }

	if res == int64(1) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(429)
	w.Write([]byte("hol up mate"))

}
