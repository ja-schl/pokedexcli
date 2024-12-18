package main

import (
	"fmt"
)

func commandHelp(cfg *Configuration, args ...string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
