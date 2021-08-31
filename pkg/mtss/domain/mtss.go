package mtss

import mtssgo "github.com/piqba/mtss-go"

type Mtsser interface {
	FetchAllFromAPI(int32) ([]mtssgo.Mtss, error)
	CreateOne(string, mtssgo.Mtss) error
}
