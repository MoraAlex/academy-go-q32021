package service

import (
	"errors"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var mockResp = `{
	"id": 19,
	"name": "rattata"
	}`

var pokResp = &entities.Pokemon{
	ID:   19,
	Name: "rattata",
}

func TestGetPokemonApi(t *testing.T) {
	testCases := []struct {
		name            string
		id              string
		err             error
		hasError        bool
		mockApiResp     string
		mockApiRespCode int
		expectedResp    *entities.Pokemon
	}{
		{
			name:            "Return a pokemon",
			id:              "2",
			err:             nil,
			hasError:        false,
			mockApiResp:     mockResp,
			mockApiRespCode: 200,
			expectedResp:    pokResp,
		},
		{
			name:            "Return error if incorrect parameter",
			id:              "a",
			err:             errors.New("Not found"),
			hasError:        true,
			mockApiResp:     "",
			mockApiRespCode: 404,
			expectedResp:    nil,
		},
	}
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	for _, tc := range testCases {
		responder := httpmock.NewStringResponder(tc.mockApiRespCode, tc.mockApiResp)
		api_url := "https://pokeapi.co/api/v2/pokemon/" + tc.id
		httpmock.RegisterResponder("GET", api_url, responder)
		s := NewGetPokemonApi(client)
		pokemons, err := s.GetPokemonApi(tc.id)
		if tc.hasError {
			assert.Error(t, err)
			assert.EqualValues(t, tc.err.Error(), err.Error())
		} else {
			assert.EqualValues(t, tc.expectedResp, pokemons)
		}
	}
}
