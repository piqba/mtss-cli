package mtss

import (
	"github.com/go-redis/redis/v8"
	"github.com/piqba/mtss-go"
)

type Mtsser interface {
	FetchAllFromAPI(int32) ([]mtss.Mtss, error)
	CreateOne(string, mtss.Mtss) error
	SendDataStream(rdb *redis.Client, key string, value map[string]interface{}) error
}
