package repository

import (
	"os"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/gocarina/gocsv"
)

type pokeRepo struct{}

//NewPokemon returns a new (pokeRepo struct {})
func NewPokemon() pokeRepo {
	return pokeRepo{}
}

//GetAll: This method returns a new([]*entities.Pokemon, error) from a CSV
func (p pokeRepo) GetAll() ([]*entities.Pokemon, error) {
	pokemonsFile, err := os.OpenFile("./utils/pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer pokemonsFile.Close()
	pokemons := []*entities.Pokemon{}
	if err := gocsv.UnmarshalFile(pokemonsFile, &pokemons); err != nil {
		return nil, err
	}
	return pokemons, nil
}
