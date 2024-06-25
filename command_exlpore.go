package main

import (
	"fmt"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/calls"
	"github.com/Groskilled/pokedex/internal/config"
)



func commandExplore(conf *config.Config, cache *cache.Cache, area string) error {
	err := calls.ExploreLocation(conf, cache, area)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
