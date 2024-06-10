package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicrodriguezval/unit-test/models"
	"github.com/nicrodriguezval/unit-test/utils"
)

var (
  ErrPokemonNotFound = errors.New("pokemon not found")
  ErrPokeapiFailure = errors.New("unexpected response from pokeapi")
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

  apiPokemon, err := GetPokemonFromPokeapi(id)
  if errors.Is(err, ErrPokemonNotFound) {
    respondwithJSON(w, http.StatusNotFound, fmt.Sprintf("pokemon with id %s not found", id))
  }
  if err != nil {
    respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error while calling pokeapi: %s", err.Error()))
  }

	parsedPokemon, err := utils.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(w, http.StatusOK, parsedPokemon)
}

func GetPokemonFromPokeapi(id string) (*models.PokeapiPokemonResponse, error) {
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
    return nil, err
	}

  if response.StatusCode == http.StatusNotFound {
    return nil, ErrPokemonNotFound
  }

	body, err := io.ReadAll(response.Body)
	if err != nil {
    return nil, err
	}

  defer response.Body.Close()

	var apiPokemon *models.PokeapiPokemonResponse

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
    return nil, err
	}

  return apiPokemon, nil
}

func respondwithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}
