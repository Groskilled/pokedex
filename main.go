package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
)

func commandHelp(conf *config.Config) error {
	fmt.Println("This is the help command")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}

func commandMap(conf *config.Config) error {
	err := calls.GetNextLocations(conf)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func commandMapB(conf *config.Config) error {
	err := calls.GetPrevLocations(conf)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func commandExit(conf *config.Config) error {
	fmt.Println("Bye !")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}

func getCommands(conf config.Config) map[string]cliCommand {
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
	commands := getCommands(conf)

	for {
		fmt.Print("pokedex > ")               // Print the prompt
		input, err := reader.ReadString('\n') // Read the input until newline
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the newline character at the end of the input
		input = input[:len(input)-1]

		if cmd, exists := commands[input]; exists {
			cmd.callback(&conf)
		}

	}
}
