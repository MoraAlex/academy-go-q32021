package usecase

import (
	"github.com/MoraAlex/academy-go-q32021/entities"
)

type pokemonRepo interface {
	GetAll() ([]*entities.Pokemon, error)
}

type getPokemonsUseCase struct {
	Repo pokemonRepo
}

//NewGetPokemons takes (repo pokemonRepo {GetAll() ([]*entities.Pokemon, error)}) as parameter
// and returns a new (getPokemonsUseCase struct {Repo pokemonRepo})
func NewGetPokemons(repo pokemonRepo) getPokemonsUseCase {
	return getPokemonsUseCase{Repo: repo}
}

//GetPokemons: This method returns ([]*entities.Pokemon, error)
func (uc getPokemonsUseCase) GetPokemons() ([]*entities.Pokemon, error) {
	pokemons, err := uc.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, err
}
