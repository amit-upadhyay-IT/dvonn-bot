package bot

import "github.com/amit-upadhyay-it/dvonn/dvonn"

// this should implement Command type
type PlayNextCommand struct {
	bot DvonnBot
}

// TODO: make this singleton with thread safety
func GetPlayNextCommand(bot DvonnBot) PlayNextCommand {
	return PlayNextCommand{bot}
}

func (playNextCmd PlayNextCommand) Execute(game *dvonn.DvonnGame) (bool, []string) {
	return playNextCmd.bot.playNextMove(game)
}
