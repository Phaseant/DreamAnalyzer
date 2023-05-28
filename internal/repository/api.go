package repository

import "context"

type ApiRepo struct {
	gpt *gptClient
}

func newApiRepo(gpt *gptClient) *ApiRepo {
	return &ApiRepo{
		gpt: gpt,
	}
}

func (api *ApiRepo) NewRequest(text string, lang int) (string, error) {
	return api.gpt.NewRequest(context.Background(), text, lang)
}
