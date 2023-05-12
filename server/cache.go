package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lucasmellof/mojangcache/model"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	redisConnection *redis.Client
	ctx             = context.Background()
)

func redisConn() *redis.Client {
	if redisConnection == nil {
		fmt.Println("Connecting to redis")
		config, err := getConfig()
		if err != nil {
			fmt.Printf("Error on get config: %s", err)
			return nil
		}
		redisConnection = redis.NewClient(&redis.Options{
			Addr:        config.Url,
			Password:    config.Password,
			Username:    config.Username,
			DB:          0,
			DialTimeout: 3,
		})
	}

	return redisConnection
}

type CacheResponse struct {
	HasCache bool                 `json:"hascache"`
	Response model.MojangResponse `json:"response"`
}

func ExistsCache(key string) int64 {
	conn := redisConn()

	return conn.Exists(ctx, key).Val()
}

func SaveValue(key string, value string) bool {
	conn := redisConn()
	conn.Set(ctx, key, value, 0)
	return true
}

func GetValue(key string) string {
	conn := redisConn()
	return conn.Get(ctx, key).Val()
}

func HasCache(key string) CacheResponse {
	hasCache := ExistsCache(key)
	value := GetValue(key)

	var response model.MojangResponse
	if hasCache == 1 {
		err := json.Unmarshal([]byte(value), &response)
		if err != nil {
			log.Println(err, "Error on unmarshal json. Key:", key)
		}
	}
	return CacheResponse{hasCache == 1, response}
}

func SaveCache(key string, response model.MojangResponse) CacheResponse {
	saved := false

	if response.Code < 500 {
		json, err := json.Marshal(response)
		if err != nil {
			log.Println(err, "Error on marshal json. Key:", key)
		}
		saved = SaveValue(key, string(json))
	}
	return CacheResponse{saved, response}
}
