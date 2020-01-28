package repository

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

type Multiplicator struct {
	Multiplicator  float64   `json:"multiplicator"`
	ExpirationTime time.Time `json:"expirationTime"`
}

func SetCache(key string, multi interface{}) bool {
	Cache.Set(key, multi, cache.NoExpiration)
	return true
}

func GetCache(key string) (Multiplicator, bool) {
	var multi Multiplicator
	var found bool
	data, found := Cache.Get(key)

	if found {
		multi = data.(Multiplicator)
	}
	return multi, found
}
