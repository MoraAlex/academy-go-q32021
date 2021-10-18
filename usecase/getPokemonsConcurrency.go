package usecase

import (
	"github.com/MoraAlex/academy-go-q32021/entities"
)

type getPokemonsConcurrencyService interface {
	GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error)
}

type pokemonsConcurrency struct {
	getPokemonsConcurrency getPokemonsConcurrencyService
}

//NewPokemonsConcurrency takes (service getPokemonsConcurrencyService interface
// {GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error)}) as parameter
// and returns a new (getPokemonsUseCase struct {Repo pokemonRepo})
func NewPokemonsConcurrency(service getPokemonsConcurrencyService) pokemonsConcurrency {
	return pokemonsConcurrency{getPokemonsConcurrency: service}
}

//GetPokemonsConcurrency: This method takes (t string, items int, ipw int) as parameters and returns ([]*entities.Pokemon, error)
func (s pokemonsConcurrency) GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error) {
	pokemons, err := s.getPokemonsConcurrency.GetPokemonsConcurrency(t, items, ipw)
	return pokemons, err
}
