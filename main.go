package main

import (
	"gogin/configs"
	"gogin/router"
	"log"
)

func main() {
	configs.Init()
	router := router.Init(&configs.Initialization{})
	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

}
