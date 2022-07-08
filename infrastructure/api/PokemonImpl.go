package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sony/gobreaker"

	"ws_amaris/domain/dto"
)

type PokemonImpl struct {
}

var circuitBreaker *gobreaker.CircuitBreaker

func InitPokemonImpl() *PokemonImpl {
	return &PokemonImpl{}
}

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= uint32(3) &&
			failureRatio >= float64(0.6)
	}

	circuitBreaker = gobreaker.NewCircuitBreaker(st)
}

func (e *PokemonImpl) PokemonById(id string) (*dto.PokemonNameDto, error) {

	var (
		name   dto.PokemonNameDto
		client http.Client
	)

	respBody, err := circuitBreaker.Execute(func() (interface{}, error) {

		url := os.Getenv("POKEMON_API") + id
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.New("error")
		}

		e.paramHeader(request)

		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		respBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("error")
		}

		return respBody, nil
	})

	if err != nil {
		return nil, err
	}

	json.Unmarshal(respBody.([]byte), &name)
	return &name, nil
}

func (e *PokemonImpl) paramHeader(request *http.Request) {
	request.Header.Set("Accepts", "application/json")
}
