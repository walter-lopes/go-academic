package repository

import (
	"time"

	. "../models"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

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
