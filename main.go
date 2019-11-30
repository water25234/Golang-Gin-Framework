package main

import (
	"./router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":3000")
}
