package pokedexclient

import (
	"encoding/json"
	"io"
	"net/http"
	cache "pokedex/pokecache"
	"time"
)

type pokemonResponse struct {
    PokemonEncounters []struct {
	Pokemon struct {
	    Name string `json:"name"`
	} `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

var pokemonCache = cache.NewCache(time.Minute)
const baseUrl = "https://pokeapi.co/api/v2/location-area/"

func GetPokemons(location string) ([]string, error) {
    url := baseUrl + location
    var err error

    data, exist := pokemonCache.Get(location)
    if !exist {
	resp, err := http.Get(url)
	if err != nil {
	    return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
	    return nil, err
	}

	pokemonCache.Add(location, data)
    }

    var r = pokemonResponse{}
    err = json.Unmarshal(data, &r)
    if err != nil {
	return nil, err
    }

    pokemons := []string{} 
    for _, pokemon := range r.PokemonEncounters {
	pokemons = append(pokemons, pokemon.Pokemon.Name)
    }

    return pokemons, nil
}
