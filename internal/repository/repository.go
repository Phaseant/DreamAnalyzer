package repository

type Api interface {
	NewRequest(text string, lang string) (string, error)
}

type Repository struct {
	Api
}

func NewRepository(gpt *gptClient) *Repository {
	return &Repository{
		Api: newApiRepo(gpt),
	}
}
