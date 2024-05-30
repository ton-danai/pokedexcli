package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationAreEncouter, err := cfg.pokeapiClient.LocationArea(name)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, data := range locationAreEncouter.PokemonEncounters {
		fmt.Printf(" - %s\n", data.Pokemon.Name)
	}

	return nil
}
