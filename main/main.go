package main

import (
	"../generator"
	"fmt"
)

func main() {
	fmt.Println("started")
	generator.GenerateNGames(5000)
	fmt.Println("finished")
}
