package main

import (
	"math/rand"
	"errors"
	"fmt"
	pokedexclient "pokedex/pokedex_client"
)

var pokemons = map[string]pokedexclient.Pokemon{}

func catchCommand(c *config, parameters []string) error {
    if len(parameters) != 1 {
	return errors.New("Please enter the name of the pokemon.")
    }

    pokemon, err := pokedexclient.GetPokemon(parameters[0])
    if err != nil {
	return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", parameters[0])
    chance := rand.Int31()
    if chance % int32(pokemon.BaseExperience) / 10 == 0 {
	pokemons[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
    } else {
	fmt.Printf("%s escaped!\n", pokemon.Name)
    }

    return nil
}
