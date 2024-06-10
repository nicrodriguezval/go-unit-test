package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
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

func TestGetPokemonFromPokeapiWithMocksSuccess(t *testing.T) {
  c := require.New(t)

  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  body, err := os.ReadFile(utils.POKEAPI_RESPONSE_PATH)
  c.NoError(err)

  id := "bulbasaur"
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

  httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(http.StatusOK, string(body)))

  pokemon, err := controller.GetPokemonFromPokeapi(id)
  c.NoError(err)

  var expectedPokemon *models.PokeapiPokemonResponse

  err = json.Unmarshal(body, &expectedPokemon)
  c.NoError(err)

  c.Equal(expectedPokemon, pokemon)
}

func TestGetPokemon(t *testing.T) {
  c := require.New(t)

  r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
  c.NoError(err)

  vars := map[string]string{
    "id": "bulbasaur",
  }

  r = mux.SetURLVars(r, vars)
  w := httptest.NewRecorder()

  controller.GetPokemon(w, r)

  var pokemon *models.Pokemon

  err = json.Unmarshal(w.Body.Bytes(), &pokemon)
  c.NoError(err)

  expectedPokemon := utils.ReadJsonFile[*models.Pokemon](t, utils.API_RESPONSE_PATH)

  c.Equal(http.StatusOK, w.Code)
  c.Equal(expectedPokemon, pokemon)
}
