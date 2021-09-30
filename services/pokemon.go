package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

type getter interface {
	GetAll() ([]model.Pokemon, error)
	GetByID(id string) (*model.Pokemon, error)
}

type pokemonRepo interface {
	getter
}

type service struct {
	repo pokemonRepo
}

func NewService(rep pokemonRepo) service {
	return service{rep}
}

// GetALlPokemons json repond to get All pokemons
func (s service) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := s.repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)
}

// GetPokemonById json repond to get a pokemon by ID
func (s service) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pokemons, err := s.repo.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)

}
