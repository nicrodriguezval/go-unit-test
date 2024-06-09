package models

type PokeApiPokemonResponse struct {
	Name        string        `json:"name"`
	PokemonType []PokemonType `json:"types"`
	Stats       []PokemonStat `json:"stats"`
	Id          int           `json:"id"`
}

type PokemonType struct {
	RefType BaseName `json:"type"`
	Slot    int      `json:"slot"`
}

type PokemonStat struct {
	Stat     BaseName `json:"stat"`
	BaseStat int      `json:"base_stat"`
}

type BaseName struct {
	Name string `json:"name"`
}
