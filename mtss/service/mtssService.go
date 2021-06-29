package mtss

import (
	domain "github.com/piqba/mtss-cli/mtss/domain"
	dto "github.com/piqba/mtss-cli/mtss/dto"
)

type MtssService interface {
	FetchAllFromAPI(string) ([]dto.MtssResponse, error)
	InsertOnDbFromAPI(string, string, int) error
}

type DefaultMtssService struct {
	repo domain.MtssRepository
}

func NewCustomerService(repository domain.MtssRepository) DefaultMtssService {
	return DefaultMtssService{repo: repository}
}
func (s DefaultMtssService) FetchAllFromAPI(url string) (response []dto.MtssResponse, err error) {
	mtssJobs, err := s.repo.FetchAllFromAPI(url)
	if err != nil {
		return nil, err
	}
	for _, job := range mtssJobs {
		response = append(response, job.ToDto())
	}
	return response, nil
}

func (s DefaultMtssService) InsertOnDbFromAPI(engine, url string, limit int) error {
	mtssJobs, err := s.repo.FetchAllFromAPI(url)
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
