package calls

import (
	"encoding/json"
	"fmt"
	"log"

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

func printLocations(results []Result) {
	for _, result := range results {
		fmt.Println(result.Name)
	}
}

func GetNextLocations(conf *config.Config, cache *cache.Cache) error {
	body := GetFromApi(conf.Next, cache)
	ans := locationAnswer{}
	err := json.Unmarshal(body, &ans)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	conf.Next = *ans.Next
	if ans.Previous != nil {
		conf.Previous = *ans.Previous
	}
	printLocations(ans.Results)
	return nil
}

func GetPrevLocations(conf *config.Config, cache *cache.Cache) error {
	body := GetFromApi(conf.Next, cache)
	ans := locationAnswer{}
	err := json.Unmarshal(body, &ans)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	conf.Next = *ans.Next
	if ans.Previous != nil {
		conf.Previous = *ans.Previous
	}
	printLocations(ans.Results)
	return nil
}

func ExploreLocation(conf *config.Config, cache *cache.Cache, area string) error {
	path := "https://pokeapi.co/api/v2/location-area/" + area
	body := GetFromApi(path, cache)
	var result PartialResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
