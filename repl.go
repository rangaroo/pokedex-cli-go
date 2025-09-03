package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var config config

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next        *string
	Previous    *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of the next 20 locations",
			callback: commandMap,
		},
	}
}
