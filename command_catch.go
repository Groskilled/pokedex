package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

func CommandCatch(conf *config.Config, cache *cache.Cache, name string, pokedex *pokemon.Pokedex) error {
	err := pokemon.CatchPokemon(conf, cache, name, pokedex)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
