package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, *cache.Cache, string, *pokemon.Pokedex) error
}

func getCommands(conf config.Config, cache *cache.Cache, area string, pokedex *pokemon.Pokedex) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world",
			callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Exlpore the designated location area",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the give Pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Retrive the information of the given pokemon",
			callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemon in the pokedex",
			callback:    CommandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	conf := config.Config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
	}
	cache := cache.NewCache(300 * time.Second)
	area := ""
	pokedex := pokemon.Pokedex{}
	commands := getCommands(conf, &cache, area, &pokedex)

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = input[:len(input)-1]
		splitted := strings.Split(input, " ")
		if len(splitted) > 1 {
			input, area = splitted[0], splitted[1]
		}
		if cmd, exists := commands[input]; exists {
			cmd.callback(&conf, &cache, area, &pokedex)
		}

	}
}
