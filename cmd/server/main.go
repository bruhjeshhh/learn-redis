package main

import(
	"net/http"
	"github.com/joho/godotenv"
	"log"
"os"
	"github.com/redis/go-redis/v9"
)
type Config struct{

}


func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

redisURL=os.Getenv("redis_url")

}