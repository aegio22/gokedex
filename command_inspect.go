package main

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("no arguments given- provide name of pokemon to inspect")
	}
	if len(args) > 1 {
		return errors.New("too many arguments given- provide only one pokemon to inspect")
	}
	p, ok := cfg.pokedex[args[0]]
	if !ok {
		return fmt.Errorf("pokemon not found in dex")
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Println("Stats: ")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.StatName, stat.BaseStatVal)
	}
	fmt.Println("Types: ")
	for _, t := range p.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
