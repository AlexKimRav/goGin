package main

import (
	"gogin/configs"
	"gogin/router"
	"log"
)

func main() {
	init := configs.Init()

	router := router.Init(init) //bug, this shit is not working
	err := router.Run("localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

}
