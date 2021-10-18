package usecase

import (
	"github.com/MoraAlex/academy-go-q32021/entities"
)

type getPokemonApiService interface {
	GetPokemonApi(id string) (*entities.Pokemon, error)
}

type updateCsvService interface {
	UpdateCsv(entities.Pokemon) ([]*entities.Pokemon, error)
}

type getPokemonUseCase struct {
	GetPokemonApiService getPokemonApiService
	UpdateCsvService     updateCsvService
}

//NewGetPokemon takes (getPS getPokemonApiService interface {GetPokemonApi(id string) (*entities.Pokemon, error)},
// csvS updateCsvService interface {UpdateCsv(entities.Pokemon) ([]*entities.Pokemon, error)}) as parameter
// and returns a new (getPokemonUseCase struct {GetPokemonApiService getPokemonApiService UpdateCsvService updateCsvService})
func NewGetPokemon(getPS getPokemonApiService, csvS updateCsvService) getPokemonUseCase {
	return getPokemonUseCase{GetPokemonApiService: getPS, UpdateCsvService: csvS}
}

//GetPokemon: This method returns (*entities.Pokemon, error) and save the result in a csv
func (uc getPokemonUseCase) GetPokemon(id string) (*entities.Pokemon, error) {
	pokemon, err := uc.GetPokemonApiService.GetPokemonApi(id)
	if err != nil {
		return nil, err
	}
	if _, err := uc.UpdateCsvService.UpdateCsv(*pokemon); err != nil && err.Error() != "That pokemons already exist" {
		return nil, err
	}
	return pokemon, err
}
