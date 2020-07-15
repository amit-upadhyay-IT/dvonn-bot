package oots

import (
	"../generator"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/amit-upadhyay-it/goutils/io"
	"log"
	"os"
	"path/filepath"
)

func ConstructTree(fileName string) {

	gameMoves := ReadFile(fileName)
	tree := GetTreeInstance()
	for _, gameMove := range gameMoves {

		tree.Insert(gameMove.Moves, gameMove.WinnerDetails)
	}
	v, _ := json.Marshal(tree)
	abs, _ := filepath.Abs("./data/model_02.json")
	f, _ := os.Create(abs)
	w := gzip.NewWriter(f)

	w.Write(v)
	//generator.AppendToFile(abs, v)

}

func ReadFile(fileName string) []generator.GameMovesWithResult {
	// TODO: the file size me be very big, so reading all at once will occupy lots of memory in RAM, so avoid doing
	// this by reading them in chunk and in do concurrent processing
	// As the current architecture of the storage file is like if we go on reading chunk by check then deserializing
	// to a type will need lots of manipulation of contents being read like trimming off some last chars, etc
	content, err := io.ReadFileBytes(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var gameMoves []generator.GameMovesWithResult
	err = json.Unmarshal(content, &gameMoves)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(gameMoves))
	return gameMoves
}
