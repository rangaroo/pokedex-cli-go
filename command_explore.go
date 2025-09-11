package main

import (
	"fmt"
)

func commandExplore(config *config, arguments []string) error {
	//TODO: Handle the case when arguments is empty
	if (arguments == nil) {
		return fmt.Errorf("Provide city name: 'explore <area_name>'")
	}

	cityName := arguments[0]
	pokemonsResp, err := config.pokeapiClient.ListPokemons(cityName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", cityName)
	fmt.Println("Found Pokemon:")

	for _, p:= range pokemonsResp.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}
	return nil
}
