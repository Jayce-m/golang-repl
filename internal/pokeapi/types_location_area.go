package pokeapi

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"` // *string type for when value is null pointer will be nil
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
