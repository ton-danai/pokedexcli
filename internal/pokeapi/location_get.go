package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (RespShallowLocationArea, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocationArea{}, err
	}

	result := RespShallowLocationArea{}
	err = json.Unmarshal(dat, &result)
	if err != nil {
		return RespShallowLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return result, nil
}
