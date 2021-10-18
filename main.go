package main

import (
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/handlers"
	"github.com/MoraAlex/academy-go-q32021/repository"
	"github.com/MoraAlex/academy-go-q32021/routes"
	"github.com/MoraAlex/academy-go-q32021/service"
	"github.com/MoraAlex/academy-go-q32021/usecase"
	"github.com/MoraAlex/academy-go-q32021/utils"

	"github.com/go-resty/resty/v2"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	client := resty.New()
	getPokemonsC := service.NewGetPokemonsConcurrency(config.PathCsvFile)
	getPokApi := service.NewGetPokemonApi(client, config.Api)
	updateCsvS := service.NewUpdateCsv(config.PathCsvFile)
	repo := repository.NewPokemon(config.PathCsvFile)
	ucGetPokemons := usecase.NewGetPokemons(repo)
	ucGetPokemon := usecase.NewGetPokemon(getPokApi, updateCsvS)
	ucGetPokemonsC := usecase.NewPokemonsConcurrency(getPokemonsC)
	h := handlers.New(ucGetPokemons, ucGetPokemon, ucGetPokemonsC)
	r := routes.New(h)
	log.Fatal(http.ListenAndServe(config.Port, r))
}
