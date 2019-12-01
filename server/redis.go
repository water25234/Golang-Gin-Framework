package server

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client = redisClient

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	ExampleClient()

	// fmt.Println(client)

	// fmt.Println(time.Second * 180)

	// client.Set("keyJustin", "Justin Value", time.Second*180).Err()

	// val, err := client.Get("keyJustin").Result()

	// fmt.Println("key", val)
}

func SetRedis(key string, value string, num int64) {

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

func ExampleClient() {

	err := client.Set("key11", "value22", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key11").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key11", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
