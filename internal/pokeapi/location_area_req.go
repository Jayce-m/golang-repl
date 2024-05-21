package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Using the client to make http requests to poke API

func (c *Client) ListLocationAreas() (LocationAreaResponse, error) {
	endpoint := "location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close() // closes the responses

	// catch bad responses
	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad Status Code")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Parse json and store in struct
	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreaResponse, nil

}
