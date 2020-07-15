package oots

import (
	"strings"
	"sync"
)

/*
 The three which will store the decisions based on games played
 */

type ModelTree struct {
	Root *GameStateNode  `json:"root"`
}

var singletonInstance *ModelTree = nil
var once sync.Once

func GetSingletonTreeInstance() *ModelTree {
	once.Do(func() {
		singletonInstance = &ModelTree{
			GetGameStateNode("r"),
			}
	})
	return singletonInstance
}

func GetTreeInstance() *ModelTree {
	return &ModelTree {
		GetGameStateNode("r"),
	}
}


/*
 Takes moves list and the score and construct tree accordingly
 */
func (tree *ModelTree) Insert(items []string, score string) {

	localRoot := tree.Root
	scoreRes := strings.Split(score, ";")
	winner := scoreRes[0]
	_ = scoreRes[1]  // TODO: use winner score and loser score to decide how better the result will be
	_ = scoreRes[2]  // loser score
	for _, item := range items {
		// from item get the player color, this on will be used in case of giving confidence
		playerColor := strings.Split(item, "|")[1]
		if val, ok := localRoot.ChildNode[item]; ok {
			// key already exists, so assign values accordingly and pass on
			if winner == playerColor {
				val.IncWinCount()
			} else if winner == "d" {  // game result was a draw
				val.IncDrawCount()
			} else {
				val.IncLoseCount()
			}
			val.IncSelectedCount()
			localRoot = val
		} else {
			// since, node isn't already present so get a new node, then assign values
			nextNode := GetGameStateNode(playerColor)
			if winner == playerColor {
				nextNode.IncWinCount()
			} else if winner == "d" {
				nextNode.IncDrawCount()
			} else {
				nextNode.IncLoseCount()
			}
			localRoot.ChildNode[item] = nextNode
			localRoot = nextNode  // assigning the just newly constructed node as the localRoot, coz next operations are supposed to be performed on it.
		}
	}
	localRoot.TerminationCount++
}

