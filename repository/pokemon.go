package repository

import (
	"os"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/gocarina/gocsv"
)

type pokeRepo struct {
	filePath string
}

//NewPokemon returns a new (pokeRepo struct {})
func NewPokemon(filePath string) pokeRepo {
	return pokeRepo{
		filePath: filePath,
	}
}

//GetAll: This method returns a new([]*entities.Pokemon, error) from a CSV
func (p pokeRepo) GetAll() ([]*entities.Pokemon, error) {
	pokemonsFile, err := os.OpenFile(p.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
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
