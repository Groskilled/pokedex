package calls

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/Groskilled/pokedex/internal/calls"
)

func GetLocations(conf *Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	return nil
}
