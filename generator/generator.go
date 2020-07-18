package generator

import (
	"encoding/json"
	"fmt"
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"log"
	"os"
	"strconv"
	"time"
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
	whiteMovesStore, blackMovesStore, currentTurnPlayer, _ := playPlacementPhase(dvonnGame, players)

	// movement phase starts
	whiteMovementMoves, blackMovementStore, _ := playMovementPhase(dvonnGame, currentTurnPlayer)

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


func generateGameForTrainingModel() GameMovesWithResult {

	players := getPlayers()
	dvonnGame := dvonn.GetDvonnGame(players, players[0])  // as first player owns white pieces
	moves := make([]string, 0)
	placementMoves := make([]string, 0)
	movementMoves := make([]string, 0)

	// do placement phase
	whiteMovesStore, blackMovesStore, currentTurnPlayer, placementMoves := playPlacementPhase(dvonnGame, players)

	// movement phase starts
	whiteMovementMoves, blackMovementStore, movementMoves := playMovementPhase(dvonnGame, currentTurnPlayer)

	whiteMovesStore = append(whiteMovesStore, whiteMovementMoves...)
	blackMovesStore = append(blackMovesStore, blackMovementStore...)

	moves = append(moves, placementMoves...)
	moves = append(moves, movementMoves...)

	winner, err := dvonnGame.GetGameWinner()
	if err != nil {
		log.Fatal(err)
	}

	if winner.GetWinnerColor() == dvonn.WINNER_WHITE {
		gameStore := GameMovesWithResult{moves, "w;"+strconv.Itoa(winner.GetWinnerScore())+";"+strconv.Itoa(winner.GetLoserScore())}
		return gameStore
	} else if winner.GetWinnerColor() == dvonn.WINNER_DRAW {
		gameStore := GameMovesWithResult{moves, "d;"+strconv.Itoa(winner.GetWinnerScore())+";"+strconv.Itoa(winner.GetLoserScore())}
		return gameStore
	} else {
		gameStore := GameMovesWithResult{moves, "b;"+strconv.Itoa(winner.GetWinnerScore())+";"+strconv.Itoa(winner.GetLoserScore())}
		return gameStore
	}
}

type GameMovesWithResult struct {
	Moves []string `json:"m"`
	WinnerDetails string `json:"win"`
}


func GenerateNGames(n int) {
	resultList := make([]GamePlayStore, 0)
	for i := 0; i < n ; i++ {
		gamePlayStore := generateGame()
		resultList = append(resultList, gamePlayStore)
	}
	gameStoreSerialized, _ := json.Marshal(resultList)
	AppendToFile("./data/100games.json", string(gameStoreSerialized)/*+"\n,"*/)
}

func GenerateNGamesForTrainingSet(n int) {
	start := time.Now()
	resultList := make([]GameMovesWithResult, 0)
	for i := 0; i < n ; i++ {
		gameMovesWithResult := generateGameForTrainingModel()
		resultList = append(resultList, gameMovesWithResult)
	}
	gameStoreSerialized, _ := json.Marshal(resultList)
	AppendToFile("./data/3games.json", string(gameStoreSerialized)/*+"\n,"*/)
	fmt.Printf("%s took\n", time.Since(start))
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


