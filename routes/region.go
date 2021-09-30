package routes

import (
	"github.com/MoraAlex/academy-go-q32021/repository"
	"github.com/MoraAlex/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

//Get handler routes
func GetterRegion(router *mux.Router) {
	regionRepo := repository.NewRegionRepo()
	s := services.NewRegionService(regionRepo)
	router.HandleFunc("/regions", s.GetAll)
	router.HandleFunc("/regions/{id}", s.GetByID)
}
