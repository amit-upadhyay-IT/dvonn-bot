package oots

type GameStateNode struct {
	ChildNode        map[string]*GameStateNode `json:"ch"` // inside map the key is the identifier for each node
	PlayerColor      string                    `json:"clr"`
	SelectedCount    int                       `json:"c"`
	WinCount         int                       `json:"w"`
	DrawCount        int                       `json:"d"`
	LoseCount        int                       `json:"l"`
	TerminationCount int                       `json:"t"`
}

func GetGameStateNode(playClr string) *GameStateNode {
	return &GameStateNode{ChildNode: make(map[string]*GameStateNode), PlayerColor:playClr, SelectedCount:1, WinCount:0, DrawCount:0, LoseCount:0}
}

func (node *GameStateNode) IncSelectedCount() {
	node.SelectedCount++
}

func (node *GameStateNode) IncWinCount() {
	node.WinCount++
}

func (node *GameStateNode) IncLoseCount() {
	node.LoseCount++
}

func (node *GameStateNode) IncDrawCount() {
	node.DrawCount++
}