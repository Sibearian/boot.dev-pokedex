package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
    reader := bufio.NewScanner(os.Stdin)
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
	    if err := cmd.callback(); err != nil {
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

type cliCommand struct {
    name string
    desc string
    callback func() error
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
    }
}
