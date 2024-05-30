package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("you must provide a pokemon name")
	}

	if pokemon, found := cfg.pokedex.Dex[args[0]]; found {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Type:")
		for _, types := range pokemon.Types {
			fmt.Printf("  - %s\n", types.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
