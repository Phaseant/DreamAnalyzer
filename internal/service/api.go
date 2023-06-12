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

// TODO Check for empty input
// TODO check lang smh...
func (api *ApiService) NewRequest(text, lang string) (string, error) {
	return api.repo.Api.NewRequest(text, lang)
}

// // We recieve lang as int from frontend, but we need to pass string to repository to translate to the original language, so we convert it here.
// func convertLang(lang int) string {
// 	switch lang {
// 	case 0:
// 		return "english"
// 	case 1:
// 		return "russian"
// 	default:
// 		return "english"
// 	}
// }
