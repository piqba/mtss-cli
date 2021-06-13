package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	mtss "github.com/PinchandoEnCuba/mtss-cli/mtss/cli"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	endpointURI := os.Getenv("ENDPOINT_MTSS")

	// redis setup
	rdb := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URI"),  // use default Addr
		Password:     os.Getenv("REDIS_PASS"), // no password set
		DB:           0,
		DialTimeout:  60 * time.Second,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		// use default DB
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)

	mtss.NewMtssCliService(rdb).FetchMtssJOBHandler(ctx, endpointURI)
}
