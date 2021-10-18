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

//NewGetPokemonsConcurrency takes (filePath string) as parameter and returns a new (getPokemonsConcurrencyService struct {filepath string})
func NewGetPokemonsConcurrency(filePath string) getPokemonsConcurrencyService {
	return getPokemonsConcurrencyService{filePath}
}

//GetPokemonsConcurrency: This method takes (t string, items int, ipw int) as parameters and returns (*entities.Pokemon, error)
//t: string. Only can be odd or even
//items: int. The number of pokemons that this method returns
//ipw: int. items that a single work can handle
func (s getPokemonsConcurrencyService) GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error) {
	jobs := make(chan []string)
	res := make(chan entities.Pokemon)
	var workersLimit int
	if items < ipw {
		workersLimit = 1
	} else {
		workersLimit = items / ipw
	}
	for i := 1; i <= workersLimit; i++ {
		go worker(i, jobs, res)
	}
	pokemonsFile, err := os.OpenFile(s.filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := csv.NewReader(pokemonsFile)
	pokemons := []*entities.Pokemon{}
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF || i == items {
			close(jobs)
			break
		}
		if record[0] == "id" {
			continue
		}
		idint, err := strconv.Atoi(record[0])
		switch t {
		case "odd":
			if idint%2 == 0 {
				jobs <- record
				pokemon := <-res
				pokemons = append(pokemons, &pokemon)
				i++
			}
		case "even":
			if idint%2 != 0 {
				jobs <- record
				pokemon := <-res
				pokemons = append(pokemons, &pokemon)
				i++
			}
		default:
			jobs <- record
			pokemon := <-res
			pokemons = append(pokemons, &pokemon)
			i++
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	defer pokemonsFile.Close()
	return pokemons, nil
}

func worker(wid int, jobs <-chan []string, result chan<- entities.Pokemon) {
	fmt.Println(wid, jobs, result)
	for j := range jobs {
		id, err := strconv.Atoi(j[0])
		if err != nil {
			fmt.Println(err)
		}
		p := entities.Pokemon{
			ID:         id,
			Name:       j[1],
			MainType:   j[2],
			SecondType: j[3],
		}
		result <- p
	}
}
