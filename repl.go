package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap map[string]cliCommand

func startRepl() {

	commandMap = map[string]cliCommand{
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
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := cleanInput(scanner.Text())
		if len(line) > 0 {
			commandName := line[0]
			commandFound := false
			for command := range commandMap {
				if command == commandName {
					err := commandMap[commandName].callback()
					if err != nil {
						fmt.Printf("Error: %v", err)
					}
					commandFound = true
				}
			}
			if !commandFound {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {

	words := strings.Fields(strings.ToLower(text))

	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()

	for command := range commandMap {
		fmt.Printf("%s: %s\n", commandMap[command].name, commandMap[command].description)
	}

	return nil
}
