package main

import (
	"./router"
)

func main() {
	router.SetupRouter()
	router.StartServer()
}
