package bot

import (
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"log"
	"math/rand"
	"time"
)

type BeginnerBot struct {}

func GetBeginnerBot() *BeginnerBot {
	return &BeginnerBot{}
}

/*
Summary: This method returns a move which can be played. This doesn't play the move itself in the game, i.e. it
		should be the responsibility of the game simulator to consume the result of this method and play a move.
		This one is the low graded bot as there is no intelligence involved in deciding the next move steps.
 How: This selects a random moves out of the possible moves and returns them

 Input: A game instance is passed for looking into the current state of the game.

 Returns: returns a tuple of boolean and selected Ids on which next move can be played

 Fun fact: sometimes this random bot will make you sweat while playing :P
*/
func (beginnerBot BeginnerBot) playNextMove(game *dvonn.DvonnGame) (bool, []string) {

	if game.IsGameOver() {
		return false, nil
	}

	// TODO: verify if this is really the right player turn and after change tense of comment to past
	currentPlayer := game.GetCurrentPlayer()

	moveIds := make([]string, 0)

	if game.GetGamePhase() == dvonn.PLACEMENT_PHASE {
		if currentPlayer.GetPlayerColor() == dvonn.WHITE {
			// get the empty fields and then choose one place among them
			selectedId := selectRandomIdForPlacementPhase(game)
			if selectedId == "" {
				return false, nil
			}
			moveIds = append(moveIds, selectedId)
		} else {
			selectedId := selectRandomIdForPlacementPhase(game)
			if selectedId == "" {
				return false, nil
			}
			moveIds = append(moveIds, selectedId)
		}
	} else {
		if currentPlayer.GetPlayerColor() == dvonn.WHITE {
			whiteCellIds := game.GetBoard().GetCellIdsByStackColor(dvonn.WHITE)
			origId, destId := selectRandomIdForOriginAndDestinationPlace(game, whiteCellIds)
			moveIds = append(moveIds, origId)
			moveIds = append(moveIds, destId)
		} else {
			blackCellIds := game.GetBoard().GetCellIdsByStackColor(dvonn.BLACK)
			origId, destId := selectRandomIdForOriginAndDestinationPlace(game, blackCellIds)
			moveIds = append(moveIds, origId)
			moveIds = append(moveIds, destId)
		}
	}

	if len(moveIds) == 0 {
		log.Fatal("no possible moves found from the bot, please look into issue as this should not have happened")
		return false, nil  // return an error also in such case as this is undesirable
	}

	return true, moveIds
}


func selectRandomIdForPlacementPhase(game *dvonn.DvonnGame) string {
	if game.IsGameOver() {
		return ""
	}
	vacantIds := make([]string, 0)
	for k, v := range game.GetBoard().GetCells() {
		// if value isn't occupied then add the key
		if v.IsEmpty() {
			vacantIds = append(vacantIds, k)
		}
	}
	if len(vacantIds) == 0 {
		return ""
	}
	// now select one among it randomly
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(vacantIds))
	return vacantIds[r]
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

