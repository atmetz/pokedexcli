package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonNames struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) SeeMap(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreas{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locations := LocationAreas{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(url, dat)

	return locations, nil
}

func (c *Client) Explore(location string) (PokemonNames, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonNames{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonNames{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonNames{}, err
	}

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonNames{}, err
	}

	pokemons := PokemonNames{}
	err = json.Unmarshal(dat, &pokemons)
	if err != nil {
		return PokemonNames{}, err
	}

	c.cache.Add(url, dat)

	return pokemons, nil

}
