package usecase

import (
	"errors"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

var mockPokemons = []*entities.Pokemon{
	{
		ID:   1,
		Name: "Test1",
	},
	{
		ID:   2,
		Name: "Test2",
	},
	{
		ID:   3,
		Name: "Test3",
	},
}

func (s serviceMock) GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error) {
	args := s.Called(t, items, ipw)
	return args[0].([]*entities.Pokemon), args.Error(1)
}

func TestGetPokemonsConcurrency(t *testing.T) {
	testCases := []struct {
		name     string
		t        string
		items    int
		ipw      int
		err      error
		hasError bool
		expResp  []*entities.Pokemon
	}{
		{
			name:     "happy path",
			t:        "odd",
			items:    3,
			ipw:      2,
			err:      nil,
			hasError: false,
			expResp:  mockPokemons,
		},
		{
			name:     "Error",
			t:        "odd",
			items:    3,
			ipw:      2,
			err:      errors.New("Test"),
			hasError: true,
			expResp:  mockPokemons,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sMock := serviceMock{}
			sMock.On("GetPokemonsConcurrency", tc.t, tc.items, tc.ipw).Return(tc.expResp, tc.err)
			uc := NewPokemonsConcurrency(sMock)
			pokemons, err := uc.GetPokemonsConcurrency(tc.t, tc.items, tc.ipw)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.items, len(pokemons))
			}
		})
	}
}
