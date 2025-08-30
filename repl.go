package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		text := cleanInput(scanner.Text())
		
		fmt.Printf("Your command was: %s\n", text[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
