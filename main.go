package main

import (
	"time"

	"github.com/ton_danai/pokedexcli/internal/pokeapi"
	"github.com/ton_danai/pokedexcli/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokeDex := pokedex.NewPokedex()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokeDex,
	}

	startRepl(cfg)
}
