package oots

import (
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"sync"
)

/*
 The three which will store the decisions based on games played
 */

type ModelTree struct {
	root *gameStateNode
}

var singletonInstance *ModelTree = nil
var once sync.Once

func GetSingletonInstance() *ModelTree {
	once.Do(func() {
		singletonInstance = &ModelTree{
			GameStateNode("0", "0", dvonn.RED),
			}
	})
	return singletonInstance
}

func GetInstance() *ModelTree {
	return &ModelTree {
		GameStateNode("0", "0", dvonn.RED),
	}
}


func (tree *ModelTree) Insert(items []string) {

}

