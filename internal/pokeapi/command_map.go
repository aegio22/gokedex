package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type mapResult struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

var nextURL *string
var prevURL *string

func commandMap() error {
	client := &http.Client{}
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if nextURL != nil && *nextURL != "" {
		url = *nextURL

	}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var locAreas mapResult
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locAreas); err != nil {
		fmt.Println(err)
	}
	if len(locAreas.Results) == 0 {
		fmt.Println("error: no locations found")
	}
	nextURL = locAreas.Next
	prevURL = locAreas.Previous
	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
