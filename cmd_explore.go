package main

import (
	"errors"
	"fmt"
	pokedexclient "pokedex/pokedex_client"
)

func exploreCommand(c *config, p []string) error {
    if len(p) != 1 {
	return errors.New("Please enter a location name like:\nexplore <location name>")
    }

    pokemons, err := pokedexclient.GetPokemons(p[0])
    if err != nil {
	return err
    }
    
    fmt.Printf("Exploring %s...\nFound Pokemon:\n", p[0])
    for _, pokemon := range pokemons {
	fmt.Printf(" - %s\n", pokemon)
    }

    return nil
}
