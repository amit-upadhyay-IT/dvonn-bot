package generator

import (
	"fmt"
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"log"
	"math/rand"
	"time"
)

func getPlayers() []dvonn.Player {

	players := make([]dvonn.Player, 0)
	player1 := dvonn.GetPlayer("amit", "upadhyay", dvonn.WHITE)
	player2 := dvonn.GetPlayer("katrina", "kaif", dvonn.BLACK)
	players = append(players, player1)
	players = append(players, player2)

	return players
}

/*
plays the moves in placement phase and returns the white player moves and black player moves respectively
 */
func playPlacementPhase(game *dvonn.DvonnGame, players []dvonn.Player) ([]string, []string, dvonn.Player) {

	cellIds := getIds()
	currentPlayer := players[0]

	whiteIds, blackIds := getWhiteAndBlackPiecesForPlacementPhase()

	whitePlacementCounter := 0
	blackPlacementCounter := 0

	whiteMovesStore := make([]string, 0)
	blackMovesStore := make([]string, 0)

	for {
		moveIds := make([]string, 0)
		if game.GetGamePhase() == dvonn.PLACEMENT_PHASE {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				moveIds = append(moveIds, cellIds[whiteIds[whitePlacementCounter]])
				whitePlacementCounter += 1
			} else {
				moveIds = append(moveIds, cellIds[blackIds[blackPlacementCounter]])
				blackPlacementCounter += 1
			}
		} else {
			break  // as for movement phase different logic will be required
		}
		moveRes := game.Move(currentPlayer, moveIds...)
		if moveRes.IsActionSuccess() {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				whiteMovesStore = append(whiteMovesStore, moveIds[0])
			} else {
				blackMovesStore = append(blackMovesStore, moveIds[0])
			}
		}
		if moveRes.IsGameOver() {
			break
		}
		if !moveRes.IsActionSuccess() {
			fmt.Println(moveRes.GetErrorCode())
			fmt.Println(moveRes.GetErrorMessage())
		}
		currentPlayer = moveRes.GetNextPlayer()
	}

	return whiteMovesStore, blackMovesStore, currentPlayer
}


func playMovementPhase(game *dvonn.DvonnGame, currentTurnPlayer dvonn.Player) ([]string, []string) {

	whiteMovesStore := make([]string, 0)
	blackMovesStore := make([]string, 0)

	for {
		moveIds := make([]string, 0)
		if game.GetGamePhase() == dvonn.MOVEMENT_PHASE {
			if currentTurnPlayer.GetPlayerColor() == dvonn.WHITE {
				// get white cells ids and select any random id which has at-least one free adjacent
				whiteCellIds := game.GetBoard().GetCellIdsByStackColor(dvonn.WHITE)
				orig, dest := selectRandomIdForOriginAndDestinationPlace(game, whiteCellIds)
				moveIds = append(moveIds, orig)
				moveIds = append(moveIds, dest)
			} else {
				// get black cells ids and select any random id which has at-least one free adjacent
				blackCells := game.GetBoard().GetCellIdsByStackColor(dvonn.BLACK)
				orig, dest := selectRandomIdForOriginAndDestinationPlace(game, blackCells)
				moveIds = append(moveIds, orig)
				moveIds = append(moveIds, dest)
			}
		}
		moveRes := game.Move(currentTurnPlayer, moveIds...)
		if moveRes.IsActionSuccess() {
			if currentTurnPlayer.GetPlayerColor() == dvonn.WHITE {
				whiteMovesStore = append(whiteMovesStore, moveIds[0] + ":" + moveIds[1])
			} else {
				blackMovesStore = append(blackMovesStore, moveIds[0] + ":" + moveIds[1])
			}
		}
		if moveRes.IsGameOver() {
			break
		}
		if !moveRes.IsActionSuccess() {
			fmt.Println(moveRes.GetErrorCode())
			fmt.Println(moveRes.GetErrorMessage())
		}
		currentTurnPlayer = moveRes.GetNextPlayer()
	}
	return whiteMovesStore, blackMovesStore
}


func getIds() []string {
	return []string{
		"a1", "a2", "a3", "b1", "b2", "b3", "b4", "c1", "c2", "c3", "c4", "c5", "d1", "d2", "d3", "d4",
		"d5", "e1", "e2", "e3", "e4", "e5", "f1", "f2", "f3", "f4", "f5", "g1", "g2", "g3", "g4", "g5", "h1", "h2",
		"h3", "h4", "h5", "i1", "i2", "i3", "i4", "i5", "j2", "j3", "j4", "j5", "k3", "k4", "k5",
	}
}

func getWhiteAndBlackPiecesForPlacementPhase() ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	whiteIds := make([]int, 0)
	for {
		r := rand.Intn(49)
		if contains(whiteIds, r) {
			continue
		} else {
			whiteIds = append(whiteIds, r)
		}
		if len(whiteIds) == 25 {
			break
		}
	}
	blackIds := shuffleItemsInList(listDifference(generateRangeItems(0, 49), whiteIds))
	return whiteIds, blackIds
}

// FYSA
func shuffleItemsInList(array []int) []int {
	for i := len(array)-1; i >= 0; i-- {
		rand.Seed(time.Now().UnixNano())
		ind := rand.Intn(len(array))
		// swap that element with last index element
		array[ind], array[i] = array[i], array[ind]
	}
	return array
}


// cellIds should be the set of ids for either white or black
func selectRandomIdForOriginAndDestinationPlace(dvonnGame *dvonn.DvonnGame, cellIds []string) (string, string) {
	originRes := ""
	destRes := ""
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(cellIds))
	loopCounter := 0
	for {
		// writing exit condition that if counter for this loop exceeds 55, either something is wrong or game is over
		if loopCounter > 55 {
			log.Fatal("game should be ended, please dont try to get origin and destination id now")
		}
		loopCounter++

		randomNode := dvonnGame.GetBoard().GetCells()[cellIds[n]]
		if randomNode.HasFreeEdge() {
			originRes = cellIds[n]
			// now see possible moves available from this cellsIds[n] id and if any is available return this id along with the destination id
			possibleMoves := dvonnGame.GetBoard().GetPossibleMoveFor(cellIds[n])
			if len(possibleMoves) > 0 {
				// selecting one
				rand.Seed(time.Now().UnixNano())
				ind := rand.Intn(len(possibleMoves))
				destRes = possibleMoves[ind].GetIdentifier()
				break
			} else {
				// as no possible move is there for cellIds[n], so we try again with a new random value of n
				rand.Seed(time.Now().UnixNano())
				n = rand.Intn(len(cellIds))
			}
		} else {
			// since no free edge node is found, try again will a new random value of n
			rand.Seed(time.Now().UnixNano())
			n = rand.Intn(len(cellIds))
		}
	}
	return originRes, destRes
}


func contains(master []int, n int) bool {
	for _, e := range master {
		if n == e {
			return true
		}
	}
	return false
}

func generateRangeItems(min, max int) []int {
	res := make([]int, 0)
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}

func listDifference(one, two []int) []int {
	res := make([]int, 0)
	for _, item1 := range one  {
		if !contains(two, item1) {
			res = append(res, item1)
		}
	}
	return res
}
