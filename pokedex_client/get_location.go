package pokedexclient

import (
	"encoding/json"
	"io"
	"net/http"
)
	
type respose struct {
	Count    int    `json:"count"`
	Next     *string	`json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Get_locations(url string) (next, prev *string, locations []string, err error) {
    resp, err := http.Get(url)
    if err != nil {
	return
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
	return
    }

    var r = respose{}
    err = json.Unmarshal(data, &r)
    if err != nil {
	return
    }

    for _, location := range r.Results {
	locations = append(locations, location.Name)
    }

    return r.Next, r.Previous, locations, nil
}
