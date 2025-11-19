package main

import (
	"fmt"
)

func CommandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, cmd := range getCommands() {
		line := fmt.Sprintf("%v: %v", cmd.name, cmd.description)
		fmt.Println(line)
	}
	return nil

}
