package routes

import (
	"github.com/MoraAlex/academy-go-q32021/repository"
	"github.com/MoraAlex/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

//Get handler routes
func GetterPokemon(router *mux.Router) {
	pokerepo := repository.NewPokemonRepo()
	s := services.NewService(pokerepo)
	router.HandleFunc("/pokemons", s.GetAllPokemons)
	router.HandleFunc("/pokemons/{id}", s.GetPokemonByID)
}
