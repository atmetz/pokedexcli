package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atmetz/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(con *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := cleanInput(scanner.Text())
		if len(line) == 0 {
			continue
		}
		commandName := line[0]
		args := []string{}
		if len(line) > 1 {
			args = line[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(con, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {

	words := strings.Fields(strings.ToLower(text))

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <Pokemon name>",
			description: "Attempt to chatch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <Pokemon name>",
			description: "See the stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See a list of all caught pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
