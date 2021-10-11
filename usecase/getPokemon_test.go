package usecase

import (
	"errors"
	"testing"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var p = &entities.Pokemon{
	Name:       "Bulbasaur",
	ID:         1,
	MainType:   "grass",
	SecondType: "Poison",
}

type MockGetPokemonApiService struct {
	mock.Mock
}

func (ms MockGetPokemonApiService) GetPokemonApi(id string) (*entities.Pokemon, error) {
	args := ms.Called(id)
	return args.Get(0).(*entities.Pokemon), args.Error(1)
}

type MockpdateCsvService struct {
	mock.Mock
}

func (ms MockpdateCsvService) UpdateCsv(p entities.Pokemon) ([]*entities.Pokemon, error) {
	args := ms.Called(p)
	return args.Get(0).([]*entities.Pokemon), args.Error(1)
}

func TestGetPokemon(t *testing.T) {
	testCases := []struct {
		name        string
		responseApi *entities.Pokemon
		responseCsv []*entities.Pokemon
		hasError    bool
		error       error
		id          string
	}{
		{
			name:        "return pokemon if id exist",
			responseApi: p,
			responseCsv: nil,
			hasError:    false,
			error:       nil,
			id:          "1",
		},
		{
			name:        "Not found if incorrect param",
			responseApi: nil,
			responseCsv: nil,
			hasError:    true,
			error:       errors.New("test error"),
			id:          "a",
		},
		{
			name:        "Empty Param",
			responseApi: nil,
			responseCsv: nil,
			hasError:    true,
			error:       errors.New("test error"),
			id:          "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGetPokemonApiService := MockGetPokemonApiService{}
			mockpdateCsvService := MockpdateCsvService{}
			mockGetPokemonApiService.On("GetPokemonApi", tc.id).Return(tc.responseApi, tc.error)
			mockpdateCsvService.On("UpdateCsv", *p).Return(tc.responseCsv, tc.error)
			uc := NewGetPokemon(mockGetPokemonApiService, mockpdateCsvService)
			pokemon, err := uc.GetPokemon(tc.id)
			if tc.hasError {
				assert.Nil(t, pokemon)
				assert.Error(t, err, "Not found")
			} else {
				assert.EqualValues(t, tc.responseApi, pokemon)
			}
		})
	}
}
