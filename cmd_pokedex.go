package main

import "fmt"

func pokedexCommand(c *config, p []string) error {
    if len(pokemons) == 0 {
	fmt.Println("You have not catched any pokemon")
    }

    fmt.Println("Your Pokedex:")
    for name := range pokemons {
	fmt.Printf(" - %s\n", name)
    }

    return nil
}
