package middleware

import (
	"fmt"
	"time"

	"github.com/water25234/Golang-Gin-Framework/server"
)

func Init() {
	fmt.Println(1234)
}

func executeThrottle() {
	fmt.Println(1234567)
	//maxAttempts := 60
	decayMinutes := 1

	time := time.Now().UTC().Format("2006-01-02 03:04:05")
	key := "test:" + time
	value := time
	decaysecond := decayMinutes * 60

	server.SetRedis(key, value, decaysecond)
}
