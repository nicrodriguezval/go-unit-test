package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicrodriguezval/unit-test/controllers"
)

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/pokemon/{id}", controller.GetPokemon).Methods("GET")

  fmt.Println("Starting server on port 8080")
  err := http.ListenAndServe(":8080", router)
  if err != nil {
    fmt.Println("Error starting server")
  }
}
