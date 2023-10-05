package main

import (
	"errors"
	"fmt"
	pokedexclient "pokedex/pokedex_client"
)

func mapCommand(c *config) error {
    if c.next == nil {
	return errors.New("You are on the last page.")
    }

    next, prev, locations, err := pokedexclient.Get_locations(*c.next)
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

func mapbCommand(c *config) error {
    if c.prev == nil {
	return errors.New("You are on the first page.")
    }

    next, prev, locations, err := pokedexclient.Get_locations(*c.prev)
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

