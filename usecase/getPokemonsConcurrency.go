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

func NewPokemonsConcurrency(service getPokemonsConcurrencyService) pokemonsConcurrency {
	return pokemonsConcurrency{getPokemonsConcurrency: service}
}

func (s pokemonsConcurrency) GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error) {
	pokemons, err := s.getPokemonsConcurrency.GetPokemonsConcurrency(t, items, ipw)
	return pokemons, err
}
