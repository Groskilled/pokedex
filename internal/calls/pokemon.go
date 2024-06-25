package calls

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
)

type pokemonAnswer struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type PartialResponse struct {
	PokemonEncounters []pokemonAnswer `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string
	Height         int
	Weight         int
	Hp             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
	Types          []string
}

type PartialResponsePokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func CatchPokemon(conf *config.Config, cache *cache.Cache, pokemon string) error {
	path := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	body := getFromApi(path, cache)
	var result PartialResponsePokemon
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if rand.Intn(10) > 2 {
		fmt.Printf("%s was caught !\n", pokemon)
		//add to pokedex
	} else {
		fmt.Printf("%s escaped !\n", pokemon)
	}
	return nil
}

func InspectPokemon(conf *config.Config, cache *cache.Cache, pokemon string) error {
	path := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	body := getFromApi(path, cache)
	var result PartialResponsePokemon
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	fmt.Println(result)
	return nil

}
