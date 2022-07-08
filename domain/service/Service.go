package service

import "ws_amaris/domain/dto"

type Service interface {
	Splitter(string) dto.SplitterDto
	PokenById(string) (*dto.PokemonNameDto, *dto.Response)
	FriendlyChains(dto.FriendlyChainsDto) string
}
