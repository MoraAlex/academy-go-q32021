package service

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/gocarina/gocsv"
)

type updateCsvService struct{}

func NewUpdateCsv() updateCsvService {
	return updateCsvService{}
}

func (s updateCsvService) UpdateCsv(p entities.Pokemon) ([]*entities.Pokemon, error) {
	pokemonsFile, err := os.OpenFile("./utils/pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// rows, err := csv.NewReader(pokemonsFile).ReadAll()
	// for _, row := range rows {
	// 	for _, row := range row {
	// 		fmt.Println(row)
	// 	}
	// }
	wr := csv.NewWriter(pokemonsFile)
	w := gocsv.NewSafeCSVWriter(wr)
	defer pokemonsFile.Close()
	pokemons := []*entities.Pokemon{}
	pokemons = append(pokemons, &p)
	if err := gocsv.MarshalCSVWithoutHeaders(pokemons, w); err != nil {
		return nil, err
	}
	return pokemons, nil
}
