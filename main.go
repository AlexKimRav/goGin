package main

import "fmt"

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"price"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	fmt.Println("Test")
}
