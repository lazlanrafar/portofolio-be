package resultsWork

import model "lazlanrafar/models"

type Service interface {
	ResultsWorkRepository() (*[]model.Work, string)
}

type service struct{
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository}
}

func (s *service) ResultsWorkRepository() (*[]model.Work, string) {
	return s.repository.ResultsWorkRepository()
}