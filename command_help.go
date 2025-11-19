package main

import (
	"fmt"
)

func commandHelp(con *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	commandMap := getCommands()

	for command := range commandMap {
		fmt.Printf("%s: %s\n", commandMap[command].name, commandMap[command].description)
	}

	return nil
}
