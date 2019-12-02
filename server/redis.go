package server

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/water25234/Golang-Gin-Framework/core/log"
)

var client *redis.Client

func init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	client = redisClient

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
}

func SetRedis(key string, value string, num int) {
	log.Info("set redis")
	err := client.Set(key, value, time.Duration(num)*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

func GetRedis(key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func GetKeys(key string) []string {
	val, err := client.Keys(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
