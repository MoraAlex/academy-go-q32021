package service

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/gocarina/gocsv"
)

type updateCsvService struct {
	FilePath string
}

//NewUpdateCsv takes (file string) as parameter and returns a new (updateCsvService struct {FilePath string})
func NewUpdateCsv(file string) updateCsvService {
	return updateCsvService{FilePath: file}
}

//UpdateCsv: This method takes (p entities.Pokemon) as parameter and returns ([]*entities.Pokemon, error)
//This method save p inside csv file int the struct. If the pokemon already exists return an error
func (s updateCsvService) UpdateCsv(p entities.Pokemon) ([]*entities.Pokemon, error) {
	pokemonsFile, err := os.OpenFile(s.FilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer pokemonsFile.Close()
	r := csv.NewReader(pokemonsFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == strconv.Itoa(p.ID) {
			return nil, errors.New("That pokemons already exist")
		}
	}
	wr := csv.NewWriter(pokemonsFile)
	w := gocsv.NewSafeCSVWriter(wr)
	pokemons := []*entities.Pokemon{}
	pokemons = append(pokemons, &p)
	if err := gocsv.MarshalCSVWithoutHeaders(pokemons, w); err != nil {
		return nil, err
	}
	return pokemons, nil
}
