package repository

import (
	"encoding/json"

	. "../models"
	"github.com/go-redis/redis"
)

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func SetRedisCache(key string, multi interface{}) bool {
	c := RedisConnect()

	cacheEntry, err := json.Marshal(multi)

	if err != nil {
		return false
	}

	err = c.Set(key, cacheEntry, 0).Err()

	if err != nil {
		panic(err)
	}

	return true
}

func GetRedisCache(key string) (Multiplicator, bool) {

	c := RedisConnect()

	val, err := c.Get(key).Result()

	var found bool
	var multi Multiplicator

	if err == redis.Nil || err != nil {
		found = false
	}

	err = json.Unmarshal([]byte(val), &multi)

	if err == nil {
		found = true
	}

	return multi, found
}
