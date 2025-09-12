package main

import (
	"fmt"
	"errors"
)

func commandPokedex(config *config, args ...string) error {
	if len(config.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemons yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
