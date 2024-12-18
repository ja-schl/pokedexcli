package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, fmt.Errorf("no name given")
	}

	url := baseURL + "/pokemon/" + name

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	var respPokemon Pokemon
	err = json.Unmarshal(dat, &respPokemon)

	return respPokemon, nil
}
