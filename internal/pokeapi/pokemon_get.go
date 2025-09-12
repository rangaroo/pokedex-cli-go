package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	
	// Check cache first
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp) 
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	ok := json.Unmarshal(dat, &pokemonResp)
	if ok != nil {
		return Pokemon{}, ok
	}
	
	c.cache.Add(url, dat)
	return pokemonResp, nil
}
