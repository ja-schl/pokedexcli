package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

func startRepl(cfg *Configuration) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}
		command, found := getCommands()[commandName]
		if !found {
			fmt.Printf("Command %q not found\n", commandName)
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("error using command %q: %v\n", commandName, err)
		}
	}
}

type Configuration struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
	Previous      *string
	Next          *string
}


type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Configuration, arg ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Exit the Pokedex",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Exit the Pokedex",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "catch specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspect specified Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list the pokemon you caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
