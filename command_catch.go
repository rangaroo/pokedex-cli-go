package main

import (
	"fmt"
	"errors"
	"math/rand"
	"math"
)
func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	probability := 3.0 / math.Sqrt(float64(pokemon.BaseExperience))

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.Float64() >= probability {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	config.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
