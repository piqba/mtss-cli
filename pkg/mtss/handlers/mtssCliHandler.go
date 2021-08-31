package mtss

import (
	"fmt"
	"log"

	service "github.com/piqba/mtss-cli/pkg/mtss/service"
)

type MtssHandler struct {
	Service service.MtssService
}

func NewMtssHandler(service service.MtssService) MtssHandler {
	return MtssHandler{
		Service: service,
	}
}

func (mh *MtssHandler) FetchAllFromAPI(limit int32) {
	data, err := mh.Service.FetchAllFromAPI(limit)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

func (mh *MtssHandler) InsertOnDbFromAPI(engine string, limit int32) {
	err := mh.Service.InsertOnDbFromAPI(engine, limit)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("OK")
}
