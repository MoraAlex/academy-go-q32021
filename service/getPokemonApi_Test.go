package service

import (
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var pokemon = entities.Pokemon{
	Name:       "Bulbasaur",
	ID:         1,
	MainType:   "grass",
	SecondType: "Poison",
}

func TestGetPokemonApiService(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       entities.Pokemon
		hasError       bool
		error          error
		id             string
	}{
		{
			name:           "return pokemon if id exist",
			expectedLength: 1,
			response:       pokemon,
			hasError:       false,
			error:          nil,
			id:             "1",
		},
		{
			name:           "Not found if incorrect param",
			expectedLength: 1,
			response:       pokemon,
			hasError:       false,
			error:          nil,
			id:             "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := resty.New()
			s := NewGetPokemonApi(client)
			pokemon, err := s.GetPokemonApi(tc.id)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.response, pokemon)
			}
		})
	}
}
