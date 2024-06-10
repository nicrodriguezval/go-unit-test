package tests

import (
	"testing"

	controller "github.com/nicrodriguezval/unit-test/controllers"
	"github.com/nicrodriguezval/unit-test/models"
	"github.com/nicrodriguezval/unit-test/utils"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeapiSuccess(t *testing.T) {
  c := require.New(t)

  pokemon, err := controller.GetPokemonFromPokeapi("bulbasaur")
  c.NoError(err)

  expectedPokemon := utils.ReadJsonFile[*models.PokeapiPokemonResponse](t, utils.POKEAPI_RESPONSE_PATH)

  c.Equal(expectedPokemon, pokemon)
}
