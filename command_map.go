package main

import (
	"errors"
	"fmt"
)

func commandMap(con *config, args ...string) error {

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

func commandMapb(con *config, args ...string) error {
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
