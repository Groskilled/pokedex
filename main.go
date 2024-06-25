package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, *cache.Cache, string) error
}

func getCommands(conf config.Config, cache *cache.Cache, area string) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Exlpore the designated location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the give Pokemon",
			callback:    commandCatch,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	conf := config.Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	cache := cache.NewCache(300)
	area := ""
	commands := getCommands(conf, cache, area)

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the newline character at the end of the input
		input = input[:len(input)-1]
		splitted := strings.Split(input, " ")
		if len(splitted) > 1 {
			input, area = splitted[0], splitted[1]
		}
		if cmd, exists := commands[input]; exists {
			cmd.callback(&conf, cache, area)
		}

	}
}
