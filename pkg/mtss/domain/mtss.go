package mtss

import "github.com/piqba/mtss-go"

type Mtsser interface {
	FetchAllFromAPI(int32) ([]mtss.Mtss, error)
	CreateOne(string, mtss.Mtss) error
}
