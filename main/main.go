package main

import (
	"../bot"
	"bufio"
	"fmt"
	"github.com/amit-upadhyay-it/dvonn/dvonn"
	"log"
	"os"
	"strings"
)

func main() {

	player1 := dvonn.GetPlayer("amit", "9148501809", dvonn.WHITE)
	player2 := dvonn.GetPlayer("bot", "1231231231", dvonn.BLACK)
	game := dvonn.GetDvonnGame([]dvonn.Player{player1, player2}, player1)

	beginnerBotCmd := bot.GetPlayNextCommand(bot.GetBeginnerBot())
	botMoveInvoker := bot.GetBotCommandInvoker(beginnerBotCmd)

	for {
		if game.IsGameOver() {
			break
		}
		if game.GetCurrentPlayer().GetPlayerColor() == dvonn.WHITE {
			// get input
			inp := inputStr()
			origId, destId := "", ""
			if strings.Contains(inp, ":") {
				origId = strings.Split(inp,":")[0]
				destId = strings.Split(inp,":")[1]
				game.Move(game.GetCurrentPlayer(), []string{origId, destId}...)
			} else {
				origId = inp
				game.Move(game.GetCurrentPlayer(), []string{origId}...)
			}
		} else {
			// bot turn
			isMovePossible, ids := botMoveInvoker.PlayAMove(game)
			fmt.Println("bot played: " + strings.Join(ids, ","))

			if isMovePossible {
				game.Move(game.GetCurrentPlayer(), ids...)
			} else {
				break
			}
		}
	}

	if game.IsGameOver() {
		winner, err := game.GetGameWinner()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(winner.GetWinnerColor())
	}
}

func inputStr() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input with : if two inputs are there ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}