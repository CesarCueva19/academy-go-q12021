package model

// PokemonsData - Pokemons info
type PokemonsData struct {
	ID     int    `json:"id,ommitempty"`
	Name   string `json:"name,ommitempty"`
	Weight int    `json:"weight,ommitempty"`
}
