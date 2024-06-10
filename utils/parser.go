package utils

import (
	"errors"

	"github.com/nicrodriguezval/unit-test/models"
)

var (
	// ErrNotFoundPokemonType occurs when the type array in pokeapi response it's not found
	ErrNotFoundPokemonType = errors.New("pokemon type array not found")
	// ErrNotFoundPokemonTypeName occurs when we found type struct but no name
	ErrNotFoundPokemonTypeName = errors.New("pokemon type name not found")
)

func ParsePokemon(apiPokemon *models.PokeapiPokemonResponse) (*models.Pokemon, error) {
	if len(apiPokemon.PokemonType) < 1 {
		return nil, ErrNotFoundPokemonType
	}

	if apiPokemon.PokemonType[0].RefType.Name == "" {
		return nil, ErrNotFoundPokemonTypeName
	}

	pokemonType := apiPokemon.PokemonType[0].RefType.Name

	abilitiesMap := map[string]int{}

	for _, stat := range apiPokemon.Stats {
		parsedAbilityName, ok := models.AllowedAbilities[stat.Stat.Name]
		if !ok {
			continue
		}

		abilitiesMap[parsedAbilityName] = stat.BaseStat
	}

	parsedPokemon := &models.Pokemon{
		Id:        apiPokemon.Id,
		Name:      apiPokemon.Name,
		Power:     pokemonType,
		Abilities: abilitiesMap,
	}

	return parsedPokemon, nil
}
