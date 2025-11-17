package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func SeeMap(locationURL string) (string, string) {

	res, err := http.Get(locationURL)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	locations := LocationAreas{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}

	return locations.Next, locations.Previous
}
