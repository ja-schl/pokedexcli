package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Configuration, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("error provide one Pokemon name")
	}
	// TODO fetch pokemon info from api
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at", pokemonName+"...")

	pokeBaseExperience := pokemon.BaseExperience
	catchChance := rand.Intn(315) + 25

	if pokeBaseExperience <= catchChance {
		// pokemon is caugth
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Println(pokemonName, "was caught!")
		return nil
	}
	fmt.Println(pokemonName, "escaped!")
	return nil
}
