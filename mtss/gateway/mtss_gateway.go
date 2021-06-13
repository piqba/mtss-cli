package mtss

import mtss "github.com/piqba/mtss-cli/mtss/models"

type MtssGateway interface {
	FetchMtssJOB(url string) []mtss.MTSS
}
type CreateStorage struct {
	MtssStorage
}

func NewMtssGateway() MtssGateway {
	return &CreateStorage{
		NewMtssStorageGateway(),
	}
}

func (c *CreateStorage) FetchMtssJOB(url string) []mtss.MTSS {
	return c.fetchMtssJOB(url)
}
