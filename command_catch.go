package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
)


func commandCatch(conf *config.Config, cache *cache.Cache, name string) error {
	err := calls.CatchPokemon(conf, cache, name)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}