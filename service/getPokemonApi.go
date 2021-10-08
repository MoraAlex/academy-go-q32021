package service

import (
	"encoding/json"
	"errors"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/go-resty/resty/v2"
)

type getPokemonApiService struct {
	Client *resty.Client
}

//NewGetPokemonApi takes (client *resty.Client) as parameter and returns a new (getPokemonApiService struct {Client *resty.Client})
func NewGetPokemonApi(client *resty.Client) getPokemonApiService {
	return getPokemonApiService{Client: client}
}

//GetPokemonApi: This method takes (id string) as parameter and returns (*entities.Pokemon, error)
func (s getPokemonApiService) GetPokemonApi(id string) (*entities.Pokemon, error) {
	resp, err := s.Client.R().
		EnableTrace().
		Get("https://pokeapi.co/api/v2/pokemon/" + id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New("Not found")
	}
	pokemon := &entities.Pokemon{}
	if err := json.Unmarshal(resp.Body(), &pokemon); err != nil {
		return nil, err
	}
	return pokemon, err
}
