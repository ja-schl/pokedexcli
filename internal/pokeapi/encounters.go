package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Explore(area string) (RespPokeEncounters, error) {
	url := baseURL + "/location-area" + "/" + area

	result := RespPokeEncounters{}
	val, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &result)
		if err != nil {
			return RespPokeEncounters{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespPokeEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokeEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokeEncounters{}, err
	}
	err = json.Unmarshal(dat, &result)
	if err != nil {
		return RespPokeEncounters{}, err
	}

	c.cache.Add(url, dat)
	return result, nil
}
