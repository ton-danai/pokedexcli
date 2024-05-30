package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	successRate := pokemon.BaseExperience * 75 / 100
	rate := rand.Intn(pokemon.BaseExperience)

	if rate >= successRate {
		cfg.pokedex.Dex[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
