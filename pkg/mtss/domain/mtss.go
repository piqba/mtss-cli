package mtss

import (
	"github.com/go-redis/redis/v8"
	mtssgo "github.com/piqba/mtss-go"
)

type Mtsser interface {
	FetchAllFromAPI(int32) ([]mtssgo.Mtss, error)
	CreateOne(string, mtssgo.Mtss) error
	SendDataStream(rdb *redis.Client, key string, value map[string]interface{}) error
}
