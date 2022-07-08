package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"ws_amaris/domain/dto"

	"github.com/sony/gobreaker"
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

		//e.paramHeader(request, headers)

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

/*
func (e *EnterpriseImpl) paramHeader(request *http.Request, headers dto.HeadersDto) {
	request.Header.Set("Accepts", "application/json")
	request.Header.Set("Authorization", "Bearer "+headers.Token)
	request.Header.Add(os.Getenv("LENGUAGE_HEADER"), headers.Language)
	request.Header.Add(os.Getenv("TENANT_HEADER"), headers.TenantId)
}
*/
