package tests

import (
	"testing"

	"github.com/nicrodriguezval/unit-test/models"
	"github.com/nicrodriguezval/unit-test/utils"
	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	response := utils.ReadJsonFile[*models.PokeapiPokemonResponse](t, utils.POKEAPI_RESPONSE_PATH)

	pokemon, err := utils.ParsePokemon(response)
	c.NoError(err)

	expectedPokemon := utils.ReadJsonFile[*models.Pokemon](t, utils.API_RESPONSE_PATH)

	c.Equal(expectedPokemon, pokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)

	response := utils.ReadJsonFile[*models.PokeapiPokemonResponse](t, utils.POKEAPI_RESPONSE_PATH)
	response.PokemonType = []models.PokemonType{}

	_, err := utils.ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(utils.ErrNotFoundPokemonType, err.Error())
}
