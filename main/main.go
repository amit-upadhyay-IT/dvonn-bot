package main

import (
	"../oots"
	"path/filepath"
)

func main() {

	abs, _ := filepath.Abs("./data/game_moves.json")
	oots.ReadFile(abs)

}

