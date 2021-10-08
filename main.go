package main

import (
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/handlers"
	"github.com/MoraAlex/academy-go-q32021/repository"
	"github.com/MoraAlex/academy-go-q32021/routes"
	"github.com/MoraAlex/academy-go-q32021/service"
	"github.com/MoraAlex/academy-go-q32021/usecase"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	getPokApi := service.NewGetPokemonApi(client)
	updateCsvS := service.NewUpdateCsv()
	repo := repository.NewPokemon()
	ucGetPokemons := usecase.NewGetPokemons(repo)
	ucGetPokemon := usecase.NewGetPokemon(getPokApi, updateCsvS)
	h := handlers.New(ucGetPokemons, ucGetPokemon)
	r := routes.New(h)
	log.Fatal(http.ListenAndServe(":8080", r))
}
