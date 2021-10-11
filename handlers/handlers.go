package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/MoraAlex/academy-go-q32021/entities"
	"github.com/gorilla/mux"
)

type getPokemonUseCase interface {
	GetPokemon(id string) (*entities.Pokemon, error)
}

type getPokemonsUseCase interface {
	GetPokemons() ([]*entities.Pokemon, error)
}

type getPokemonsConcurrencyUseCase interface {
	GetPokemonsConcurrency(t string, items int, ipw int) ([]*entities.Pokemon, error)
}

type handlers struct {
	GetAllPokemonsUseCase         getPokemonsUseCase
	GetPokemonApiUseCase          getPokemonUseCase
	GetPokemonsConcurrencyUseCase getPokemonsConcurrencyUseCase
}

//New takes (GetAllPokemonsUseCase getPokemonsUseCase interface {GetPokemon(id string) (*entities.Pokemon, error)},
// GetPokemonApiUseCase getPokemonUseCase interface {GetPokemons() ([]*entities.Pokemon, error)}) as parameters
// and returns a new (handlers struct {GetAllPokemonsUseCase getPokemonsUseCase GetPokemonApiUseCase getPokemonUseCase})
func New(ucpokemons getPokemonsUseCase, ucpokemon getPokemonUseCase, usconcurrency getPokemonsConcurrencyUseCase) handlers {
	return handlers{GetAllPokemonsUseCase: ucpokemons, GetPokemonApiUseCase: ucpokemon, GetPokemonsConcurrencyUseCase: usconcurrency}
}

//GetAllPokemons: This method is the handler for pokemons
func (h handlers) GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := h.GetAllPokemonsUseCase.GetPokemons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :(" + err.Error()})
	}
	json.NewEncoder(w).Encode(pokemons)
}

//GetPokemon: This method is the handler for /pokemon/{id}
func (h handlers) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "ID not found"})
		return
	}
	matched, err := regexp.Match("^[0-9]*$", []byte(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :(", "error": err.Error()})
		return
	}
	if !matched {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad request: ID is not valid"})
		return
	}
	pokemons, err := h.GetPokemonApiUseCase.GetPokemon(id)
	if err != nil {
		if err.Error() == "Not found" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "Pokemons Not Found", "error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :(", "error": err.Error()})
		return

	}
	json.NewEncoder(w).Encode(pokemons)
	return
}

func (h handlers) GetPokemonsConcurrency(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	t, ok := vars["type"]
	items := vars["items"]
	ipw := vars["items_per_workers"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "type param not defined"})
		return
	}
	if len(t) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid number of arguments to type param"})
		return
	}
	if _, ok := vars["items"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "items param not defined"})
		return
	}
	if len(items) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid number of arguments to items param"})
		return
	}
	if _, ok := vars["items_per_workers"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "items_per_workers param not defined"})
		return
	}
	if len(ipw) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid number of arguments to items_per_workers  param"})
		return
	}
	if t[0] != "odd" && t[0] != "even" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "type param not defined"})
		return
	}
	itemsint, err := strconv.Atoi(items[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid type on param items"})
		return
	}
	ipwint, err := strconv.Atoi(ipw[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid type on param items_per_work "})
		return
	}
	pokemons, err := h.GetPokemonsConcurrencyUseCase.GetPokemonsConcurrency(strings.Join(t, ""), itemsint, ipwint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Ooops!! :(", "error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(pokemons)
	return
}
