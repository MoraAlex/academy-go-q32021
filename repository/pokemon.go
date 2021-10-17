package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/MoraAlex/academy-go-q32021/entities"
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
	reader := csv.NewReader(pokemonsFile)
	i := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			i++
			continue
		}
		idint, err := strconv.Atoi(row[0])
		pokemon := entities.Pokemon{
			ID:         idint,
			Name:       row[1],
			MainType:   row[2],
			SecondType: row[3],
		}
		pokemons = append(pokemons, &pokemon)
		i++
	}
	return pokemons, nil
}
