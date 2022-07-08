package api

import "ws_amaris/domain/dto"

type Pokemon interface {
	PokemonById(string) (*dto.PokemonNameDto, error)
}
