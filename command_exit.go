package main

import (
	"fmt"
	"os"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
)


func commandExit(conf *config.Config, cache *cache.Cache, area string) error {
	fmt.Println("Bye !")
	os.Exit(0)
	return nil
}