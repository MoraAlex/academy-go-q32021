package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/gorilla/mux"
)

type getPokemonUseCase interface {
	GetPokemon(id string) (*entities.Pokemon, error)
}

type getPokemonsUseCase interface {
	GetPokemons() ([]*entities.Pokemon, error)
}

type handlers struct {
	GetAllPokemonsUseCase getPokemonsUseCase
	GetPokemonApiUseCase  getPokemonUseCase
}

//New takes (GetAllPokemonsUseCase getPokemonsUseCase interface {GetPokemon(id string) (*entities.Pokemon, error)},
// GetPokemonApiUseCase getPokemonUseCase interface {GetPokemons() ([]*entities.Pokemon, error)}) as parameters
// and returns a new (handlers struct {GetAllPokemonsUseCase getPokemonsUseCase GetPokemonApiUseCase getPokemonUseCase})
func New(ucpokemons getPokemonsUseCase, ucpokemon getPokemonUseCase) handlers {
	return handlers{GetAllPokemonsUseCase: ucpokemons, GetPokemonApiUseCase: ucpokemon}
}

//GetAllPokemons: This method is the handler for pokemons
func (h handlers) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := h.GetAllPokemonsUseCase.GetPokemons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :("})
	}
	json.NewEncoder(w).Encode(pokemons)
}

//GetPokemon: This method is the handler for /pokemon/{id}
func (h handlers) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if ok {
		matched, err := regexp.Match("^[0-9]*$", []byte(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :("})
		}
		if !matched {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Bad request: ID is not valid"})
		} else {
			pokemons, err := h.GetPokemonApiUseCase.GetPokemon(id)
			if err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(pokemons)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "ID not found"})
	}
}
