package mtss

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	mtss "github.com/piqba/mtss-cli/mtss/models"
)

type MtssStorage interface {
	fetchMtssJOB(url string) []mtss.MTSS
}
type MtssService struct {
}

func (s *MtssService) fetchMtssJOB(url string) []mtss.MTSS {
	var mtssArray []mtss.MTSS

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &mtssArray)
	if err != nil {
		log.Fatal("Unmarshal", err)
	}

	return mtssArray
}

func NewMtssStorageGateway() MtssStorage {
	return &MtssService{}
}
