package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	mtss "github.com/piqba/mtss-cli/mtss/cli"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var (
	source = "mtss"
	batch  = false
)

var ctx = context.Background()

func init() {
	flag.StringVar(&source, "source", source, "--source <mtss> Data source to ingest on DB")
	flag.BoolVar(&batch, "batch", batch, "--batch <true> For fetch daily jobs or one by one jobs")
	flag.Parse()
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

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

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis down")
		// kill app
	}

	// CLI logic
	if source == "mtss" {
		if batch {
			mtss.NewMtssCliService(rdb).InsertJobsOneByOneOnRedis(ctx, endpointURI)
		} else {
			mtss.NewMtssCliService(rdb).InsertJobsDailyBatchOnRedis(ctx, endpointURI)
		}
	}
}
