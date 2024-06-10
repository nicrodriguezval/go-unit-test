package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/nicrodriguezval/unit-test/models"
	"github.com/nicrodriguezval/unit-test/utils"
	"github.com/stretchr/testify/require"
)

func TestParserPokemon(t *testing.T) {
	c := require.New(t)

	body, err := os.ReadFile("../utils/samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)

	pokemon, err := utils.ParsePokemon(response)
  c.NoError(err)

	body, err = os.ReadFile("../utils/samples/api_response.json")
	c.NoError(err)

	var expectedPokemon *models.Pokemon

	err = json.Unmarshal(body, &expectedPokemon)
	c.NoError(err)

  c.Equal(expectedPokemon, pokemon)
}
