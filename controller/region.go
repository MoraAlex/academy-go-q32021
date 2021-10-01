package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/model"
	"github.com/gorilla/mux"
)

type getterRegion interface {
	GetAll() ([]model.Region, error)
	GetByID(id string) (model.Region, error)
}

type regionRepo interface {
	getterRegion
}

type regionService struct {
	repo regionRepo
}

func NewRegionService(repo regionRepo) regionService {
	return regionService{repo}
}

func (rs regionService) GetAll(w http.ResponseWriter, r *http.Request) {
	b, err := rs.repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(b)
}

func (rs regionService) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	b, err := rs.repo.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(b)
}
