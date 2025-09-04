package main

import (
	"errors"
	"fmt"
)

func commandMap(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *config) error {
	if config.prevLocationsURL == nil {
		return errors.New("Can't go back")
	}
	
	locationsResp, err := config.pokeapiClient.ListLocations(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationsResp.Next
	config.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
