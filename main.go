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

const filePath = "./utils/pokemon.csv"

func main() {
	client := resty.New()
	getPokApi := service.NewGetPokemonApi(client)
	updateCsvS := service.NewUpdateCsv(filePath)
	repo := repository.NewPokemon(filePath)
	ucGetPokemons := usecase.NewGetPokemons(repo)
	ucGetPokemon := usecase.NewGetPokemon(getPokApi, updateCsvS)
	h := handlers.New(ucGetPokemons, ucGetPokemon)
	r := routes.New(h)
	log.Fatal(http.ListenAndServe(":8080", r))
}
