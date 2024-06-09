package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type Pokemon struct {
	Abilities map[string]int `json:"Abilities"`
	Name      string         `json:"Name"`
	Power     string         `json:"Type"`
	Id        int            `json:"Id"`
}

var Abilities = map[string]int{
	"Hp":      0,
	"Attack":  0,
	"Defense": 0,
	"Speed":   0,
}

var AllowedAbilities = map[string]string{
	"hp":      "Hp",
	"attack":  "Attack",
	"defense": "Defense",
	"speed":   "Speed",
}
