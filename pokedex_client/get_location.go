package pokedexclient

import (
	"encoding/json"
	"io"
	"net/http"
	cache "pokedex/pokecache"
	"time"
)
	
type locationResponse struct {
    Count    int    `json:"count"`
    Next     *string	`json:"next"`
    Previous *string    `json:"previous"`
    Results  []struct {
	    Name string `json:"name"`
	    URL  string `json:"url"`
    } `json:"results"`
}

var locationCache = cache.NewCache(time.Minute)

func GetLocations(url string) (next, prev *string, locations []string, err error) {
    data, exist := locationCache.Get(url)

    if !exist {
	resp, error := http.Get(url)
	if err != nil {
	    return nil, nil, nil, error
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
	    return nil, nil, nil, error
	}

	locationCache.Add(url, data)
    }

    var r = locationResponse{}
    err = json.Unmarshal(data, &r)
    if err != nil {
	return
    }

    for _, location := range r.Results {
	locations = append(locations, location.Name)
    }

    return r.Next, r.Previous, locations, nil
}
