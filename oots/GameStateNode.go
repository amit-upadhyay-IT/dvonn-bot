package oots

import "github.com/amit-upadhyay-it/dvonn/dvonn"

type gameStateNode struct {
	srcId string
	destId string
	playerColor dvonn.ChipColor
	selectedCount int
	winCount int
	drawCount int
	loseCount int
}

func GameStateNode(srcId, destId string, playClr dvonn.ChipColor) *gameStateNode {
	return &gameStateNode{srcId:srcId, destId:destId, playerColor:playClr, selectedCount:1, winCount:0, drawCount:0, loseCount:0}
}

func (node *gameStateNode) IncSelectedCount() {
	node.selectedCount++
}

func (node *gameStateNode) IncWinCount() {
	node.winCount++
}

func (node *gameStateNode) IncLoseCount() {
	node.loseCount++
}

func (node *gameStateNode) IncDrawCount() {
	node.drawCount++
}