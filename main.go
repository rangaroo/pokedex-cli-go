package main

import(
	"time"

	"github.com/rangaroo/pokedex-cli-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 2 * time.Second)

	config := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
