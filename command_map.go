package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
)

func commandMap(conf *config.Config, cache *cache.Cache, area string) error {
	err := calls.GetNextLocations(conf, cache)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func commandMapB(conf *config.Config, cache *cache.Cache, area string) error {
	err := calls.GetPrevLocations(conf, cache)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
