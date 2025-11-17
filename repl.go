package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// Command registry
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func startREPL() {

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
		words := cleanInput(scanner.Text())
		commandName := words[0]
		cmd, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue

		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}

}
