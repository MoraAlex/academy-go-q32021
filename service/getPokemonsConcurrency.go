package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/MoraAlex/academy-go-q32021/entities"
)

type getPokemonsConcurrencyService struct {
	filepath string
}

func NewGetPokemonsConcurrency(filePath string) getPokemonsConcurrencyService {
	return getPokemonsConcurrencyService{filePath}
}

func (s getPokemonsConcurrencyService) GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error) {
	jobs := make(chan []string, 100)
	res := make(chan entities.Pokemon, 100)
	go worker(jobs, res)
	pokemonsFile, err := os.OpenFile(s.filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := csv.NewReader(pokemonsFile)
	pokemons := []*entities.Pokemon{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			close(jobs)
			break
		}
		if record[0] == "id" {
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		jobs <- record
		pokemon := <-res
		pokemons = append(pokemons, &pokemon)
	}
	defer pokemonsFile.Close()
	return pokemons, nil
}

func worker(jobs <-chan []string, result chan<- entities.Pokemon) {
	for j := range jobs {
		pid, err := strconv.Atoi(j[0])
		if err != nil {
			fmt.Println(err)
		}
		p := entities.Pokemon{
			ID:         pid,
			Name:       j[1],
			MainType:   j[2],
			SecondType: j[3],
		}
		result <- p
	}
}
