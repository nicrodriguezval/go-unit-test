// Description: Helper functions for testing.

package utils

import (
	"encoding/json"
	"os"

	"github.com/stretchr/testify/require"
)

const (
  POKEAPI_RESPONSE_PATH = "../utils/samples/pokeapi_response.json"
  API_RESPONSE_PATH = "../utils/samples/api_response.json"
)

func ReadJsonFile[T any, Testing require.TestingT](t Testing, path string) T {
	c := require.New(t)

	body, err := os.ReadFile(path)
	c.NoError(err)

	var response T

	err = json.Unmarshal(body, &response)
	c.NoError(err)

	return response
}
