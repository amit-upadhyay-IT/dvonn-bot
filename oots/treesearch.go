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
	//v, _ := json.Marshal(tree)
	abs, _ := filepath.Abs("./data/model_500000games_opt.zip")
	f, _ := os.Create(abs)
	w := gzip.NewWriter(f)
	//
	GetMaxNestedStructureCount(tree)
	fmt.Println(maxCount)

	confidenceTree := GetConfidenceTree()
	confidenceTree.ConstructConfidenceTree(tree)
	v, _ := json.Marshal(confidenceTree)
	fmt.Println("confidence tree constructed")
	w.Write(v)
	abs, _ = filepath.Abs("./data/model_500000games_opt.json")
	generator.AppendToFile(abs, string(v))

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


func GetMaxNestedStructureCount(tree *ModelTree) {

	_GetMaxNestedStructureCount(tree.Root, 0)
}

var maxCount int = 0

func _GetMaxNestedStructureCount(root *GameStateNode, count int) {

	if root.ChildNode == nil || len(root.ChildNode) == 0 {
		if count > maxCount {
			fmt.Println("count:", count)
			//fmt.Println("upadhyay")
			maxCount = count
		}
		return
	}
	for _, v := range root.ChildNode {
		_GetMaxNestedStructureCount(v, count+1)
		//fmt.Println("amit")
	}
}