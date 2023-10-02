package main

import (
	"fmt"
)

func helpCommand() error {
    commands := getCommands()

    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()

    for _, cmd := range commands {
	fmt.Printf("%s: %s\n", cmd.name, cmd.desc)
    }

    fmt.Println()

    return nil
}
