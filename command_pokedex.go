package main

import (
	"errors"
	"fmt"
)

func commandPokedex(con *config, args ...string) error {
	if len(con.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")

	for s := range con.caughtPokemon {
		fmt.Printf(" - %v\n", s)
	}

	return nil
}
