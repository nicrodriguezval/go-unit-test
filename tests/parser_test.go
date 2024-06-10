package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/nicrodriguezval/unit-test/models"
	"github.com/nicrodriguezval/unit-test/utils"
	"github.com/stretchr/testify/require"
)

func readJsonFile[T any](t *testing.T, path string) T {
  c := require.New(t)

  body, err := os.ReadFile(path)
  c.NoError(err)

  var response T

  err = json.Unmarshal(body, &response)
  c.NoError(err)

  return response
}

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

  response := readJsonFile[models.PokeApiPokemonResponse](t, "../utils/samples/pokeapi_response.json")

	pokemon, err := utils.ParsePokemon(response)
  c.NoError(err)

  expectedPokemon := readJsonFile[*models.Pokemon](t, "../utils/samples/api_response.json")

  c.Equal(expectedPokemon, pokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
  c := require.New(t)

  response := readJsonFile[models.PokeApiPokemonResponse](t, "../utils/samples/pokeapi_response.json")
  response.PokemonType = []models.PokemonType{}

  _, err := utils.ParsePokemon(response)
  c.NotNil(err)
  c.EqualError(utils.ErrNotFoundPokemonType, err.Error())
}

