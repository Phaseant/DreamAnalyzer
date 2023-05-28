package service

import "github.com/Phaseant/DreamAnalyzer/internal/repository"

type ApiService struct {
	repo *repository.Repository
}

func newApiService(repo *repository.Repository) *ApiService {
	return &ApiService{
		repo: repo,
	}
}

// TODO Clean request from Dan prefix. Maybe check for correctness, etc...
func (api *ApiService) NewRequest(text string, lang int) (string, error) {
	return api.repo.Api.NewRequest(text, lang)
}
