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

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])

	if err != nil {
		return err
	}

	rate := rand.Intn(pokemon.BaseExperience)

	if rate <= 40 {
		cfg.pokedex.Dex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
