package main

import "fmt"

func commandPokedex(cfg *Configuration, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _,v := range cfg.caughtPokemon {
		fmt.Println(" - ", v.Name)
	}
	return nil
}
