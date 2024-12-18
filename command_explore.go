package main

import (
	"fmt"
)

func commandExplore(cfg *Configuration, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide an area name")
	}
	result, err := cfg.pokeapiClient.Explore(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...", args[0])
	fmt.Println("Found Pokemon:")
	for _, p := range result.PokemonEncounters {
		fmt.Println("-", p.Pokemon.Name)
	}

	return nil
}
