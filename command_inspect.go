package main

import (
	"errors"
	"fmt"
)

func commandInspect(con *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]

	_, exists := con.caughtPokemon[name]

	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", con.caughtPokemon[name].Name)
	fmt.Printf("Height: %v\n", con.caughtPokemon[name].Height)
	fmt.Printf("Weight: %v\n", con.caughtPokemon[name].Weight)
	fmt.Println("Stats:")
	for s := range con.caughtPokemon[name].Stats {
		fmt.Printf("  -%v: %v\n", con.caughtPokemon[name].Stats[s].Stat.Name, con.caughtPokemon[name].Stats[s].BaseStat)
	}
	fmt.Println("Types:")
	for t := range con.caughtPokemon[name].Types {
		fmt.Printf("  - %v\n", con.caughtPokemon[name].Types[t].Type.Name)
	}

	return nil
}
