package main

import "fmt"

func commandInspect(cfg *Configuration, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("give one name of a pokemon")
	}
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println(name, "has not been caught yet!")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
	return nil
}
