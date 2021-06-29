package mtss

import (
	"fmt"
	"log"

	service "github.com/piqba/mtss-cli/mtss/service"
)

type MtssHandler struct {
	Service service.MtssService
}

func (mh *MtssHandler) FetchAllFromAPI(url string) {
	data, err := mh.Service.FetchAllFromAPI(url)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

func (mh *MtssHandler) InsertOnDbFromAPI(engine, url string, limit int) {
	err := mh.Service.InsertOnDbFromAPI(engine, url, limit)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("OK")
}
