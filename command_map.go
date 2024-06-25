package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

func CommandMap(conf *config.Config, cache *cache.Cache, area string, pokedex *pokemon.Pokedex) error {
	err := calls.GetNextLocations(conf, cache)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func CommandMapB(conf *config.Config, cache *cache.Cache, area string, pokedex *pokemon.Pokedex) error {
	err := calls.GetPrevLocations(conf, cache)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
