package pokedexclient

import (
	"encoding/json"
	"io"
	"net/http"
	cache "pokedex/pokecache"
	"time"
)

type pokemonsResponse struct {
    PokemonEncounters []struct {
	Pokemon struct {
	    Name string `json:"name"`
	} `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

var pokemonsCache = cache.NewCache(time.Minute)
const baseUrlPokemons = "https://pokeapi.co/api/v2/location-area/"

func GetPokemons(location string) ([]string, error) {
    url := baseUrlPokemons + location
    var err error

    data, exist := pokemonsCache.Get(location)
    if !exist {
	resp, err := http.Get(url)
	if err != nil {
	    return nil, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
	    return nil, err
	}

	pokemonsCache.Add(location, data)
    }

    var r = pokemonsResponse{}
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
