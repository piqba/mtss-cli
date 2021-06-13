package mtss

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	gateway "github.com/piqba/mtss-cli/mtss/gateway"
	models "github.com/piqba/mtss-cli/mtss/models"
)

var signature = "mtss:job"

type MtssCliService struct {
	gtw gateway.MtssGateway
	rdb *redis.Client
}

func NewMtssCliService(rdb *redis.Client) *MtssCliService {
	return &MtssCliService{
		gtw: gateway.NewMtssGateway(),
		rdb: rdb,
	}
}

// InsertJobsOneByOneOnRedis it`s used for insert simple hash on redis
// redis command [hset key field1 value1 ...]
func (s *MtssCliService) InsertJobsOneByOneOnRedis(ctx context.Context, url string) {

	mtssJobs := s.gtw.FetchMtssJOB(url)
	for _, job := range mtssJobs {
		go func(job models.MTSS) {
			key := fmt.Sprintf("%s:%d", signature, job.ID)
			value, _ := job.ToMAP()
			if _, err := s.rdb.HSet(ctx, key, value).Result(); err != nil {
				log.Fatal("create: redis error: %w", err)
			}
			s.rdb.Expire(ctx, key, 24*time.Hour)
		}(job)
	}
	log.Println("Ingested all jobs: ", len(mtssJobs))
}

func (s *MtssCliService) InsertJobsDailyBatchOnRedis(ctx context.Context, url string) {
	currentTime := time.Now()
	date := strings.Split(currentTime.String(), " ")[0]
	keySufix := strings.ReplaceAll(
		date,
		"-",
		"",
	)
	key := fmt.Sprintf("%s:%s", signature, keySufix)

	mtssJobs := s.gtw.FetchMtssJOB(url)
	bytes, err := json.Marshal(mtssJobs)
	if err != nil {
		log.Fatalln(err)
	}
	dcache := models.DailyCache{
		ID:    keySufix,
		Jobs:  string(bytes),
		Count: len(mtssJobs),
	}

	value, _ := dcache.ToMAP()

	if _, err := s.rdb.HSet(ctx, key, value).Result(); err != nil {
		log.Fatal("create: redis error: %w", err)
	}
	s.rdb.Expire(ctx, key, 24*time.Hour)

}
