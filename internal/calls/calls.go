package calls

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAnswer struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

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
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
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

func getFromApi(path string, cache *cache.Cache) []byte {
	var body []byte
	cached := false
	body, cached = cache.Get(path)
	if !cached {
		res, err := http.Get(path)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		cache.Add(path, body)
	}
	return body
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
