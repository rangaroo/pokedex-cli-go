package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMap(config *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area"

	if config.Next != nil {
		baseURL = *config.Next
	}

	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
		return err
	}

	locationArea := locationArea{}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	fmt.Println()

	config.Next = locationArea.Next
	config.Previous = locationArea.Previous

	return nil
}

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
