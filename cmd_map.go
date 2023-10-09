package main

import (
	"errors"
	"fmt"
	pokedexclient "pokedex/pokedex_client"
)

func mapCommand(c *config, p []string) error {
    if c.next == nil {
	return errors.New("You are on the last page.")
    }

    next, prev, locations, err := pokedexclient.GetLocations(*c.next)
    if err != nil {
	return err
    }

    c.next = next
    c.prev = prev

    for _, location := range locations {
	fmt.Println(location)
    }

    return nil
}

func mapbCommand(c *config, p []string) error {
    if c.prev == nil {
	return errors.New("You are on the first page.")
    }

    next, prev, locations, err := pokedexclient.GetLocations(*c.prev)
    if err != nil {
	return err
    }

    c.next = next
    c.prev = prev

    for _, location := range locations {
	fmt.Println(location)
    }

    return nil
}

