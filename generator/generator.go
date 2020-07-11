package generator

import (
	"encoding/json"
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"log"
	"os"
)


/*
 * Generator idea:
	Each dvonnGame instrument corresponds to a game played
	The idea is that play games and record the playing moves with the end result
	First I start with placement phase and then the movement phase.
	For placement phase, I generate a list of random moves for white player and black player eg: whitePlaces := {"d4","g1","f3","h5","b1","e1","f1","g2","g3","d1","h3","f4"... etc} and similarly I select
		Ids for the black player.
	NOTE that for placement phase which can only have 25 places and black can have 24 places(as white turn is first)

	For the movement phase:
	Select a random move(i.e. origin Id, destination id) for white player or black player depending on the current
	player turn. Play the turn as and when required. And on game end store the movements and result.
 */
func generateGame() GamePlayStore {

	players := getPlayers()
	dvonnGame := dvonn.GetDvonnGame(players, players[0])  // as first player owns white pieces

	// do placement phase
	whiteMovesStore, blackMovesStore, currentTurnPlayer := playPlacementPhase(dvonnGame, players)

	// movement phase starts
	whiteMovementMoves, blackMovementStore := playMovementPhase(dvonnGame, currentTurnPlayer)

	whiteMovesStore = append(whiteMovesStore, whiteMovementMoves...)
	blackMovesStore = append(blackMovesStore, blackMovementStore...)

	winner, err := dvonnGame.GetGameWinner()
	if err != nil {
		log.Fatal(err)
	}



	gameStore := GamePlayStore{whiteMovesStore, blackMovesStore, winner}
	return gameStore

}

type GamePlayStore struct {
	WhiteMoves []string
	BlackMoves []string
	WinnerRes *dvonn.MatchResult
}


func GenerateNGames(n int) {
	resultList := make([]GamePlayStore, 0)
	for i := 0; i < n ; i++ {
		gamePlayStore := generateGame()
		resultList = append(resultList, gamePlayStore)
	}
	gameStoreSerialized, _ := json.Marshal(resultList)
	AppendToFile("./data/game_moves.json", string(gameStoreSerialized)/*+"\n,"*/)
}

func AppendToFile(filename, value string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		file.Close()
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(value)); err != nil {
		return err
	}
	return nil
}


