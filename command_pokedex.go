package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

func CommandPokedex(conf *config.Config, cache *cache.Cache, name string, pokedex *pokemon.Pokedex) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range pokedex.Pokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
