package pokeapi

import (
	"errors"
	"fmt"
	"math/rand"
)

func (c *Client) TryCatch(pm *Pokemon) (bool, error) {
	if pm == nil {
		return false, errors.New("no pokemon given for catch try")
	}

	base := pm.BaseExperience
	const (
		minExp       = 30
		maxExp       = 650
		minExpChance = 0.95
		maxExpChance = 0.05
	)

	if base < minExp {
		base = minExp
	}
	if base > maxExp {
		base = maxExp
	}

	expRange := float64(maxExp - minExp)
	chanceRange := minExpChance - maxExpChance
	catchChance := minExpChance - (float64(base-minExp)/expRange)*chanceRange

	roll := rand.Float64()

	fmt.Printf("Base Experience: %d\n", base)
	fmt.Printf("Catch Chance: %.2f%%\n", catchChance*100)
	fmt.Printf("Roll: %.2f\n", roll)

	if roll < catchChance {
		fmt.Printf("%s was caught!\n", pm.Name)
		return true, nil
	}

	fmt.Printf("%s escaped!\n", pm.Name)
	return false, nil

}
