package middlewares

import (
	"fmt"
)

func Init() {
	fmt.Println(1234)
}

func executeThrottle() {
	//maxAttempts := 60
	//decayMinutes := 1

	//time := time.Now().UTC().Format("2006-01-02 03:04:05")
	//key := "test:" + time
	//value := time
	//decaysecond := decayMinutes * 60

	//server.SetRedis(key, value, decaysecond)
}
