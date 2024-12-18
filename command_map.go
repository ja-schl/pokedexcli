package main

import (
	"fmt"
)

func commandMapF(cfg *Configuration, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	return nil
}

func commandMapB(cfg *Configuration, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := cfg.pokeapiClient.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	return nil
}
