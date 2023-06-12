package service

import "github.com/Phaseant/DreamAnalyzer/internal/repository"

type Api interface {
	NewRequest(text, lang string) (string, error)
}

type Service struct {
	Api
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Api: newApiService(repos),
	}
}
