package main

import (
	"fmt"
	"os"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
	"github.com/Groskilled/pokedex/internal/pokemon"
)

func CommandExit(conf *config.Config, cache *cache.Cache, area string, pokedex *pokemon.Pokedex) error {
	fmt.Println("Bye !")
	os.Exit(0)
	return nil
}
