package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler interface {
	GetAllPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemon(w http.ResponseWriter, r *http.Request)
	GetPokemonsConcurrency(w http.ResponseWriter, r *http.Request)
}

//NewGetPokemonApi takes (handler handler interface {GetAllPokemons(w http.ResponseWriter,
// r *http.Request) GetPokemon(w http.ResponseWriter, r *http.Request)}) as parameter
// and returns a new (*mux.Router)
func New(handler handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemons", handler.GetAllPokemons).Methods(http.MethodGet)
	r.HandleFunc("/pokemons/{id}", handler.GetPokemon).Methods(http.MethodGet)
	r.HandleFunc("/pokemons-concurrency", handler.GetPokemonsConcurrency).Methods(http.MethodGet)
	return r
}
