package main

import(
	"time"

	"github.com/rangaroo/pokexe-cli-go/internal"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
