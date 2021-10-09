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

func NewUpdateCsv(file string) updateCsvService {
	return updateCsvService{FilePath: file}
}

func (s updateCsvService) UpdateCsv(p entities.Pokemon) ([]*entities.Pokemon, error) {
	pokemonsFile, err := os.OpenFile(s.FilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
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
	defer pokemonsFile.Close()
	pokemons := []*entities.Pokemon{}
	pokemons = append(pokemons, &p)
	if err := gocsv.MarshalCSVWithoutHeaders(pokemons, w); err != nil {
		return nil, err
	}
	return pokemons, nil
}
