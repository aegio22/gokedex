package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func CommandBMap() error {
	if prevURL == nil || *prevURL == "" {
		return fmt.Errorf("you're on the first page")
	}
	client := &http.Client{}
	url := *prevURL

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
