package server

import (
	"encoding/json"
	"github.com/dominik-najberg/gogoapps/client"
	"github.com/patrickmn/go-cache"
	"log"
	"strings"
	"time"
)

var c *cache.Cache

func Cache() client.Adapter {
	return func(f client.Fetcher) client.Fetcher {
		return func(url string) (string, error) {
			log.Println("check cache:", url)
			if jsonStr, found := c.Get(url); found {
				log.Println("found in cache:", url)
				return jsonStr.(string), nil
			}

			jsonStr, err := f(url)
			if err != nil {
				return "", err
			}

			if validateResponseCodeSuccessful(jsonStr) {
				log.Printf("update cache: %v -> %s", url, jsonStr)
				c.Set(url, jsonStr, cache.DefaultExpiration)
			}

			return jsonStr, nil
		}
	}
}

// Do not cache pages if error occurred
func validateResponseCodeSuccessful(s string) bool {
	type CodeCheck struct {
		Cod int `json:"cod"`
	}
	a := &CodeCheck{}
	if err := json.NewDecoder(strings.NewReader(s)).Decode(a); err != nil {
		return false
	}

	if a.Cod != 200 {
		return false
	}

	return true
}

func InitCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}