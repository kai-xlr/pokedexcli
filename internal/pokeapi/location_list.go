package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	//Url
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	//Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//Response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return RespShallowLocations{}, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	//Extract data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//Unmarshal data
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
