package oots

import (
	"../generator"
	"encoding/json"
	"fmt"
	"github.com/amit-upadhyay-it/goutils/io"
	"log"
)

func ConstructTree(fileName string) {

	gamesPlayed := ReadFile(fileName)


}

func ReadFile(fileName string) []generator.GamePlayStore {
	// TODO: the file size me be very big, so reading all at once will occupy lots of memory in RAM, so avoid doing
	// this by reading them in chunk and in do concurrent processing
	// As the current architecture of the storage file is like if we go on reading chunk by check then deserializing
	// to a type will need lots of manipulation of contents being read like trimming off some last chars, etc
	content, err := io.ReadFileBytes(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var gameMoves []generator.GamePlayStore
	err = json.Unmarshal(content, &gameMoves)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(gameMoves))
	return gameMoves
}
