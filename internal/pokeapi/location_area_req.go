package pokeapi

import "net/http"

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

}
