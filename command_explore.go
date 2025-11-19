package main

import (
	"errors"
	"fmt"
)

func commandExplore(con *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	location := args[0]
	fmt.Printf("Exploring: %s...", location)

	pokemonsResp, err := con.pokeapiClient.Explore(location)
	if err != nil {
		return err
	}

	fmt.Println("\nFound Pokemon:")
	for _, pok := range pokemonsResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}

	return nil
}
