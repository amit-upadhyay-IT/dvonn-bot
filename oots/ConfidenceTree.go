package oots

import "math"

type ConfidenceTree struct {
	Root *ConfidenceNode
}

func GetConfidenceTree() *ConfidenceTree {
	return &ConfidenceTree{Root:GetConfidenceNode()}
}


type ConfidenceNode struct {
	Child      map[string]*ConfidenceNode `json:"c"`
	SuitableId string                     `json:"s,omitempty"`
}

func GetConfidenceNode() *ConfidenceNode {
	return &ConfidenceNode{Child:make(map[string]*ConfidenceNode), SuitableId:""}
}

/*
 Constructs the Confidence Tree by traversing over the Model Tree
 */
func (confTree *ConfidenceTree) ConstructConfidenceTree(modelTree *ModelTree) {
	traverseModelTreeAndConstructConfidenceTree(modelTree.Root, confTree.Root)
}

/*
 Traverse the Model Tree and construct the ConfidenceTree by adding new nodes in Confidence Tree as and when required
 */
func traverseModelTreeAndConstructConfidenceTree(modelRoot *GameStateNode, confRoot *ConfidenceNode) {
	if modelRoot == nil || len(modelRoot.ChildNode) == 0 {
		// since at this point, the confidence node should also have no child nodes as the game moves are over,
		// so we don't assign any preferredCellId to this one
		return
	}
	// TODO using sync pool which help improve the performance
	for k, v := range modelRoot.ChildNode {
		// get a confidence node with that Id
		// create an empty child map and assign to the new node
		childMap := make(map[string]*ConfidenceNode)
		preferredCellId := GetNodeWithMaxConfidence(modelRoot.ChildNode, modelRoot.SelectedCount)

		newConfNode := GetConfidenceNode()
		newConfNode.Child = childMap
		newConfNode.SuitableId = preferredCellId
		confRoot.Child[k] = newConfNode
		traverseModelTreeAndConstructConfidenceTree(v, confRoot.Child[k])
	}
}

/*
 Returns the id of the child node which have the most confidence
 */
func GetNodeWithMaxConfidence(gameChildNodes map[string]*GameStateNode, visitCount int) string {

	maxConfidenceId := ""
	maxConf := 0.0
	// calculate the confidence for each of the child nodes
	for id, v := range gameChildNodes {
		conf := calculateConfidence(v.WinCount, v.SelectedCount, visitCount)

		// keeping the strict less coz, if confidence is coming out as zero, default at there should
		// be another logic for selecting the preferred move and it should be handled by the client bot
		if conf > maxConf {
			maxConf = conf
			maxConfidenceId = id
		}
	}
	return maxConfidenceId
}

func calculateConfidence(vi, ni, visitCount int) float64 {
	return float64(vi) + math.Sqrt(math.Log(float64(visitCount)/float64(ni)))
}
