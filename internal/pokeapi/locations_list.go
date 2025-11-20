package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type mapResult struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []locArea `json:"results"`
}

type locArea struct {
	Name               string `json:"name"`
	URL                string `json:"url"`
	Pokemon_Encounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Stat struct {
			StatId   int    `json:"id"`
			StatName string `json:"name"`
		} `json:"stat"`
		BaseStatVal int `json:"base_stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) ListLocations(pageURL *string) (mapResult, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	var locAreas mapResult
	if val, ok := c.cache.Get(url); !ok {
		fmt.Println("cache not used")
		res, err := c.httpClient.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		c.cache.Add(url, body)

		if err := json.Unmarshal(body, &locAreas); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("cache used")
		if err := json.Unmarshal(val, &locAreas); err != nil {
			fmt.Println(err)
		}
	}

	if len(locAreas.Results) == 0 {
		fmt.Println("error: no locations found")
	}

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return locAreas, nil
}
