package main

import (
	"fmt"
	"log"
	"os"

	"actions/internal/client/redis"
	"actions/internal/div"
	"actions/internal/sum"
)

// this should obv be an API, but I'll keep it simple since this is a demo
func main() {
	fmt.Printf("4 + 1 = %d\n", sum.Sum(4, 1))
	fmt.Printf("4 / 1 = %d\n", div.CustomDivide(4, 1))

	// all base configuration is done in the `main` pkg
	_, err := redis.NewRedisClient(os.Getenv(redis.RedisHostEnvVar))
	if err != nil {
		log.Fatalf("could not connect to redis")
	}
}
