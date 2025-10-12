package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(LocationName string) (Location, error) {
	url := baseURL + "/location-area/" + LocationName

	if val, ok := c.cache.Get(url); ok {
		locationResponse := Location{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return Location{}, err
		}
		return locationResponse, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return Location{}, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResponse := Location{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)
	return locationResponse, nil
}
