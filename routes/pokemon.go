package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	GetAllPokemons(w http.ResponseWriter, r *http.Request)
	GetPokemonByID(w http.ResponseWriter, r *http.Request)
}

func New(handler Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemons", handler.GetAllPokemons).Methods(http.MethodGet)
	r.HandleFunc("/pokemons/{id}", handler.GetPokemonByID).Methods(http.MethodGet)
	return r
}

//Get handler routes
// func GetterPokemon(router *mux.Router) {
// 	pokerepo := repository.NewPokemonRepo()
// 	s := services.NewService(pokerepo)
// 	router.HandleFunc("/pokemons", s.GetAllPokemons)
// 	router.HandleFunc("/pokemons/{id}", s.GetPokemonByID)
// }
