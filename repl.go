package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/atmetz/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(con)
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

func commandExit(con *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(con *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	commandMap := getCommands()

	for command := range commandMap {
		fmt.Printf("%s: %s\n", commandMap[command].name, commandMap[command].description)
	}

	return nil
}

func commandMap(con *config) error {

	locationsResp, err := con.pokeapiClient.SeeMap(con.next)
	if err != nil {
		return err
	}

	con.next = locationsResp.Next
	con.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(con *config) error {
	if con.previous == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := con.pokeapiClient.SeeMap(con.previous)
	if err != nil {
		return err
	}

	con.next = locationsResp.Next
	con.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
