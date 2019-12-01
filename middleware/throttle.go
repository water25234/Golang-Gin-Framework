package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/water25234/Golang-Gin-Framework/core/log"
	core "github.com/water25234/Golang-Gin-Framework/core/server"
	"github.com/water25234/Golang-Gin-Framework/server"
)

func init() {
	fmt.Println(1234)
}

func ExecuteThrottle() gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println(555)

		log.Info("Execute Throttle")
		//maxAttempts := 60
		decayMinutes := 1

		time := time.Now().UTC().Format("2006-01-02 03:04:05")
		key := core.GetServerConfig().IpAddress + ":" + time
		value := time
		decaysecond := decayMinutes * 60

		server.SetRedis(key, value, decaysecond)

	}
}
