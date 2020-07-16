package main

import (
	"../oots"
	//"../generator"
	"path/filepath"
)

func main() {

	abs, _ := filepath.Abs("./data/5000games.json")
	oots.ConstructTree(abs)

	//generator.GenerateNGamesForTrainingSet(1)

}

