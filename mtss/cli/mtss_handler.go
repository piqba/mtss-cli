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

func (s *MtssCliService) FetchMtssJOBHandler(ctx context.Context, url string) {
	currentTime := time.Now()
	keyField := strings.Split(currentTime.String(), " ")[0]
	mtssKey := fmt.Sprintf("mtss:%s", keyField)

	result, err := s.rdb.HGet(ctx, mtssKey, keyField).Result()
	if err != nil && err != redis.Nil {
		log.Fatal("find: redis error: %w", err)
	}
	if result == "" {

		mtssJobs := s.gtw.FetchMtssJOB(url)
		dcache := models.DailyCache{
			ID:    keyField,
			Jobs:  mtssJobs,
			Count: len(mtssJobs),
		}
		bytes, err := json.Marshal(dcache)
		if err != nil {
			log.Fatalln(err)
		}
		if _, err := s.rdb.HSetNX(ctx, mtssKey, keyField, string(bytes)).Result(); err != nil {
			log.Fatal("create: redis error: %w", err)
		}
		s.rdb.Expire(ctx, mtssKey, 24*time.Hour)
	} else {

		dcache := models.DailyCache{}
		if err := dcache.UnmarshalBinary([]byte(result)); err != nil {
			log.Fatal("find: unmarshal error: %w", err)
		}
		fmt.Println(dcache.Count)
	}

}
