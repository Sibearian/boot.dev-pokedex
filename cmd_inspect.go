package main

import (
	"errors"
	"fmt"
)

func inspectCommand(c *config, parameter []string) error {
    if len(parameter) != 1 {
	return errors.New("Please type the name of the pokemon")
    }

    pokemon, exist := pokemons[parameter[0]]
    if !exist {
	fmt.Println("you have not caught that pokemon")
	return nil
    }

    fmt.Println("Name: ", pokemon.Name)
    fmt.Println("Height: ", pokemon.Height)
    fmt.Println("Weight: ", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
	fmt.Printf("  -%s: %d\n", stat.Name, stat.BaseStat)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
	fmt.Printf("  -%s\n", t.Name)
    }

    return nil
}
