package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
)

func commandHelp() error {
	fmt.Println("This is the help command")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}

func commandMap(conf *Config) error {
	locations := calls.GetLocations()
	fmt.Println(locations)
	return nil
}

func commandExit() error {
	fmt.Println("Bye !")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
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
	commands := getCommands()

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
			cmd.callback()
		}

	}
}
