package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Configuration struct {
	Url      *string
	Previous *string
	Next     *string
}

func (c *Client) GetLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	result := RespShallowLocations{}
	entry, found := c.cache.Get(url)
	if found {
		err := json.Unmarshal(entry, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(dat, &result)
	if err != nil {
		return result, err
	}

	c.cache.Add(url, dat)
	return result, nil
}
