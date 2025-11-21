package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(pokeName *string) (Pokemon, error) {

	if pokeName == nil {
		return Pokemon{}, errors.New("no pokemon name provided")
	}

	url := baseURL + "/pokemon/" + *pokeName

	var pokemon *Pokemon
	if val, ok := c.cache.Get(url); !ok {
		fmt.Println("cache not used")
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error fetching data: %w", err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error reading response body: %w", err)
		}

		if res.StatusCode == http.StatusNotFound {
			return Pokemon{}, fmt.Errorf("pokemon '%s' not found", *pokeName)
		}
		c.cache.Add(url, body)

		if err := json.Unmarshal(body, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling pokemon data: %w", err)
		}
	} else {
		fmt.Println("cache used")
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling cached data: %w", err)
		}
	}

	return *pokemon, nil

}
