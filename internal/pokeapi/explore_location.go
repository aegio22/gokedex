package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c *Client) ExploreLocation(locName *string) (locArea, error) {

	if locName == nil {
		return locArea{}, errors.New("no location name provided")
	}
	url := baseURL + "/location-area/" + *locName

	var loc *locArea
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

		if err := json.Unmarshal(body, &loc); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("cache used")
		if err := json.Unmarshal(val, &loc); err != nil {
			fmt.Println(err)
		}
	}

	return *loc, nil

}
