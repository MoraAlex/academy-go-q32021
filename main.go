package main

import (
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/controller"
	"github.com/MoraAlex/academy-go-q32021/routes"
)

func main() {
	c := controller.New(service)
	r := routes.New(c)
	log.Fatal(http.ListenAndServe(":8080", r))
}
