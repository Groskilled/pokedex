package calls

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Groskilled/pokedex/internal/cache"
	"github.com/Groskilled/pokedex/internal/config"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Answer struct {
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

func getLocations(path string, cache *cache.Cache) Answer {
	var body []byte
	var err error
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
		fmt.Println("before Add")
		cache.Add(path, body)
		fmt.Println("Add done")
	}
	ans := Answer{}
	err = json.Unmarshal(body, &ans)
	if err != nil {
		log.Fatalf("error unmarshalling JSON: %v", err)
	}
	return ans
}

func GetNextLocations(conf *config.Config, cache *cache.Cache) error {
	ans := getLocations(conf.Next, cache)
	conf.Next = *ans.Next
	if ans.Previous != nil {
		conf.Previous = *ans.Previous
	}
	printLocations(ans.Results)
	return nil
}

func GetPrevLocations(conf *config.Config, cache *cache.Cache) error {
	ans := getLocations(conf.Previous, cache)
	conf.Next = *ans.Next
	if ans.Previous != nil {
		conf.Previous = *ans.Previous
	}
	printLocations(ans.Results)
	return nil
}
