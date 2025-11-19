package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(con *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonResp, err := con.pokeapiClient.PokemonName(name)
	if err != nil {
		return err
	}

	catchChance := rand.Intn(pokemonResp.BaseExperience)

	if catchChance < 40 {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		con.caughtPokemon[pokemonResp.Name] = pokemonResp
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonResp.Name)

	return nil
}
