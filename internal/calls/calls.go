package calls

import (
	"io"
	"log"
	"net/http"

	"github.com/Groskilled/pokedex/internal/cache"
)

func GetFromApi(path string, cache *cache.Cache) []byte {
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
