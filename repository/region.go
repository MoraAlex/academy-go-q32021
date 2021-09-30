package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/model"
)

type regionRepo struct{}

//return a regionRepo struct
func NewRegionRepo() regionRepo {
	return regionRepo{}
}

//return all regions from external API
//Docs: https://pokeapi.co/docs/v2#regions
func (r regionRepo) GetAll() ([]model.Region, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/region/")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	apiBody := model.ApiPokemonResponse{}
	if err := json.Unmarshal(b, &apiBody); err != nil {
		log.Fatal(err)
	}
	regions := apiBody.Results
	resp.Body.Close()

	return regions, err
}

//return a region by id from external API
//Docs: https://pokeapi.co/docs/v2#regions
func (r regionRepo) GetByID(id string) (model.Region, error) {
	resp, err := http.Get(`https://pokeapi.co/api/v2/region/` + id)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	region := model.Region{}
	if err := json.Unmarshal(b, &region); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	return region, err
}
