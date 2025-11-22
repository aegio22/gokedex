package main

import (
	"errors"
	"fmt"
)

func CommandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments given- provide name of pokemon to catch")
	}
	if len(args) > 1 {
		return errors.New("too many arguments given- provide only one pokemon to catch")
	}

	pokemon, err := cfg.pokeapiClient.FetchPokemon(&args[0])
	if err != nil {
		return err
	}

	if _, exists := cfg.pokedex[pokemon.Name]; exists {
		return errors.New("pokemon already in dex")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	catch, err := cfg.pokeapiClient.TryCatch(&pokemon)
	if err != nil {
		return err
	}
	if catch {
		cfg.pokedex[pokemon.Name] = pokemon
	}
	return nil

}
