package pokedexclient

import (
    "encoding/json"
    "io"
    "net/http"
    cache "pokedex/pokecache"
    "time"
)

type pokemonResponce struct {
    Name   string `json:"name"`
    Height int    `json:"height"`
    Weight int    `json:"weight"`
    BaseExperience int `json:"base_experience"`
    Stats  []struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
	    Name string `json:"name"`
	} `json:"stat"`
    } `json:"stats"`
    Types []struct {
	Slot int `json:"slot"`
	Type struct {
	    Name string `json:"name"`
	} `json:"type"`
    } `json:"types"`
}

type Pokemon struct {
    Name string
    Height int
    Weight int
    BaseExperience int
    Stats []PokeStats
    Types []PokeType    
}

type PokeStats struct {
    Name string
    BaseStat int
    Effort int
}

type PokeType struct {
    Name string
}

var pokemonCache = cache.NewCache(time.Minute)
const baseUrlPokemon = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemon(name string) (poke Pokemon, err error) {
    url := baseUrlPokemon + name

    data, exist := pokemonCache.Get(url)
    if !exist {
	resp, err := http.Get(baseUrlPokemon + name)
	if err != nil {
	    return poke, err
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
	    return poke, err
	}

	pokemonCache.Add(url, data)
    }

    var r = pokemonResponce{}
    err = json.Unmarshal(data, &r)
    if err != nil {
	return poke, err
    }

    poke.Name = r.Name
    poke.Height = r.Height
    poke.Weight = r.Weight
    poke.BaseExperience = r.BaseExperience
    
    for _, stat := range r.Stats {
	pStat := PokeStats{}
	pStat.BaseStat = stat.BaseStat
	pStat.Name = stat.Stat.Name
	pStat.Effort = stat.Effort
	poke.Stats = append(poke.Stats, pStat)
    }

    for _, rType := range r.Types {
	pType := PokeType{}
	pType.Name = rType.Type.Name
	poke.Types = append(poke.Types, pType)
    }

    return poke, nil
}
