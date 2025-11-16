package main

import (
	"strings"
)

func cleanInput(text string) []string {

	//lower := strings.ToLower(text)
	words := strings.Fields(strings.ToLower(text))

	return words
}
