package main

import "../generator"

func main() {

	//abs, _ := filepath.Abs("./data/game_moves.json")
	//oots.ReadFile(abs)

	generator.GenerateNGamesForTrainingSet(5000)

}

