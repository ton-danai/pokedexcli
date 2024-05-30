package pokedex

import "github.com/ton_danai/pokedexcli/internal/pokeapi"

type Pokedex struct {
	Dex map[string]pokeapi.Pokemon
}

func NewPokedex() Pokedex {
	return Pokedex{
		Dex: map[string]pokeapi.Pokemon{},
	}
}
