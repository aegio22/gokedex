package main

import (
	"fmt"
)

func CommandPokedex(cfg *config, args []string) error {
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("no pokemon found in dex")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
