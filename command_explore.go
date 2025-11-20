package main

import (
	"errors"
	"fmt"
)

func CommandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments given- provide location to explore")
	}
	if len(args) > 1 {
		return errors.New("too many arguments given- provide only one location to explore")
	}
	fmt.Printf("Exploring %v\n", args[0])
	loc, err := cfg.pokeapiClient.ExploreLocation(&args[0])
	if err != nil {
		return err
	}

	if len(loc.Pokemon_Encounters) == 0 {
		return errors.New("no pokemon encounters found")
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range loc.Pokemon_Encounters {
		line := "- " + pokemon.Pokemon.Name
		fmt.Println(line)
	}

	return nil
}
