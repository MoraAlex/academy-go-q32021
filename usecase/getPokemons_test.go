package usecase

import (
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pokemons = []*entities.Pokemon{
	{
		Name:       "test1",
		ID:         1,
		MainType:   "grass",
		SecondType: "Poison",
	},
	{
		Name:       "test2",
		ID:         2,
		MainType:   "grass",
		SecondType: "Poison",
	},
	{
		Name:       "test3",
		ID:         3,
		MainType:   "grass",
		SecondType: "Poison",
	},
}

type MockRepo struct {
	mock.Mock
}

func (mr MockRepo) GetAll() ([]*entities.Pokemon, error) {
	args := mr.Called()
	return args.Get(0).([]*entities.Pokemon), args.Error(1)
}

func TestGetPokemons(t *testing.T) {
	testCases := []struct {
		name     string
		response []*entities.Pokemon
		hasError bool
		error    error
		id       string
	}{
		{
			name:     "return pokemon if id exist",
			response: pokemons,
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := MockRepo{}
			mockRepo.On("GetAll").Return(tc.response, tc.error)
			uc := NewGetPokemons(mockRepo)
			pokemons, err := uc.GetPokemons()
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.response, pokemons)
			}
		})
	}
}
