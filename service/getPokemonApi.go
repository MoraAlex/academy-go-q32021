package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/entities"

	"github.com/go-resty/resty/v2"
)

type getPokemonApiService struct {
	Client *resty.Client
	Api    string
}

//NewGetPokemonApi takes (client *resty.Client) as parameter and returns a new (getPokemonApiService struct {Client *resty.Client})
func NewGetPokemonApi(client *resty.Client, api string) getPokemonApiService {
	return getPokemonApiService{Client: client, Api: api}
}

//GetPokemonApi: This method takes (id string) as parameter and returns (*entities.Pokemon, error)
func (s getPokemonApiService) GetPokemonApi(id string) (*entities.Pokemon, error) {
	resp, err := s.Client.R().
		EnableTrace().
		Get(s.Api + id)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("Not found")
	}
	pokemon := &entities.Pokemon{}
	if err := json.Unmarshal(resp.Body(), &pokemon); err != nil {
		return nil, err
	}
	return pokemon, err
}
