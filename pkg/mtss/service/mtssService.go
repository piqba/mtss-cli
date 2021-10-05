package mtss

import (
	domain "github.com/piqba/mtss-cli/pkg/mtss/domain"
	"github.com/piqba/mtss-go"
	mtssgo "github.com/piqba/mtss-go"
)

type MtssService interface {
	FetchAllFromAPI(limit int32) ([]mtssgo.Mtss, error)
	InsertOnDbFromAPI(string, int32) error
	GetMtssJobs() ([]mtss.Mtss, error)
}

type DefaultMtssService struct {
	repo domain.MtssRepository
}

func NewCustomerService(repository domain.MtssRepository) DefaultMtssService {
	return DefaultMtssService{repo: repository}
}
func (s DefaultMtssService) FetchAllFromAPI(limit int32) ([]mtssgo.Mtss, error) {
	mtssJobs, err := s.repo.FetchAllFromAPI(limit)
	if err != nil {
		return nil, err
	}
	return mtssJobs, nil
}

func (s DefaultMtssService) InsertOnDbFromAPI(engine string, limit int32) error {
	mtssJobs, err := s.repo.FetchAllFromAPI(limit)
	if err != nil {
		return err
	}
	for _, job := range mtssJobs[0:limit] {
		err := s.repo.CreateOne(engine, job)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s DefaultMtssService) GetMtssJobs() ([]mtss.Mtss, error) {
	mtssJobs, err := s.repo.GetMtssJobs()
	if err != nil {
		return nil, err
	}
	return mtssJobs, nil
}
