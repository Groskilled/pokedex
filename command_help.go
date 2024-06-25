package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

func CommandHelp(conf *config.Config, cache *cache.Cache, area string, pokedex *pokemon.Pokedex) error {
	fmt.Println("This is the help command")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}
