package service

import (
	"net/http"
	"strings"

	"ws_amaris/domain/dto"
	"ws_amaris/infrastructure/api"
)

type ServiceImpl struct {
	PokemonApi api.Pokemon
}

func InitServiceImpl() *ServiceImpl {
	return &ServiceImpl{
		PokemonApi: api.InitPokemonImpl(),
	}
}

func (r *ServiceImpl) Splitter(split string) dto.SplitterDto {

	var (
		splitter dto.SplitterDto
	)

	return splitter.SplitterOrdering(split)
}

func (r *ServiceImpl) PokenById(id string) (*dto.PokemonNameDto, *dto.Response) {

	name, err := r.PokemonApi.PokemonById(id)

	if err != nil {
		return nil, &dto.Response{
			Status:      http.StatusBadRequest,
			Description: err.Error(),
		}
	}

	return name, nil
}

func (r *ServiceImpl) FriendlyChains(text dto.FriendlyChainsDto) string {

	var (
		wo = text.TextOne
		ws = text.TextTwo
	)

	if len(wo) != len(ws) {
		return text.Message(false)
	}

	for i := 0; i < len(wo); i++ {
		u := wo[0 : i+1]
		v := wo[i+1:]

		if strings.Compare(ws, v+u) == 0 {
			return text.Message(true)
		}
	}

	return text.Message(false)
}
