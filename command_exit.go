package main

import (
	"fmt"
	"os"
)

func commandExit(con *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
