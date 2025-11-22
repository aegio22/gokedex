package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aegio22/gokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Prints a list of 20 location areas",
			callback:    CommandMapf,
		},
		"bmap": {
			name:        "bmap",
			description: "Prints a list of the previous 20 location areas",
			callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Print a list of all Pokémon located in a given location",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:	"inspect",
			description: "Get attribute values for given pokemon",
			callback: CommandInspect,
		},

	}
}

func StartREPL(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		if scanner.Err() != nil {
			err := fmt.Errorf("error encountered: %v", scanner.Err())
			fmt.Println(err)
			continue
		}
		if scanner.Text() == "" {
			fmt.Println("no text in the scanner")
			continue
		}

		var args []string
		words := cleanInput(scanner.Text())
		commandName := words[0]

		if len(words) > 1 {
			args = words[1:]
		} else {
			args = []string{}
		}

		cmd, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue

		} else {
			err := cmd.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}

}
