package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

func startRepl() {
    reader := bufio.NewScanner(os.Stdin)
    var c config = config{
	next: &url,
	prev: nil,
    }
    for {
        fmt.Printf("pokedex > ")
        reader.Scan()

	words := cleanInput(reader.Text())
	if len(words) == 0 {
	    continue
	}

	commandName := words[0]
	cmd, exist := getCommands()[commandName]
	if exist {
	    if err := cmd.callback(&c, words[1:]); err != nil {
		fmt.Println(err)
	    }
	} else {
	    fmt.Printf("Unknown command\n")
	}
    }
}

func cleanInput(text string) []string {
    line := strings.ToLower(text)
    words := strings.Fields(line)
    return words 
}

type config struct {
    next *string
    prev *string
}

type cliCommand struct {
    name string
    desc string
    callback func(c *config, p []string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
	"help" : {
	    name: "help",
	    desc: "Displays a help message",
	    callback: helpCommand,
	},
	"exit" : {
	    name: "exit",
	    desc: "Exit the Pokedex",
	    callback: exitCommand,
	},
	"map" : {
	    name: "map",
	    desc: "Get the next page of locations",
	    callback: mapCommand,
	},
	"mapb" : {
	    name: "mapb",
	    desc: "Get the previous page of locations",
	    callback: mapbCommand,
	},
	"explore" : {
	    name: "explore",
	    desc: "Explore a given location and get pokemons available",
	    callback: exploreCommand,
	},
	"catch" : {
	    name: "catch",
	    desc: "Try to catch a pokemon",
	    callback: catchCommand,
	},
    }
}
