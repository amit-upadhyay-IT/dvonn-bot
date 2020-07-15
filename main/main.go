package main

import (
	"../oots"
	"path/filepath"
)

func main() {

	abs, _ := filepath.Abs("./data/games_01.json")
	oots.ConstructTree(abs)

	//generator.GenerateNGamesForTrainingSet(5000)

}

