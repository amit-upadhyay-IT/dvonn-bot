package main

import (
	"../oots"
	"path/filepath"

	//"path/filepath"

)

func main() {

	abs, _ := filepath.Abs("./data/400000games_optimize.json")
	oots.ConstructTree(abs)

	//generator.GenerateNGamesForTrainingSet(400000)

}

