package bot

// this should implement Command type
type PlayNextCommand struct {
	bot DvonnBot
}

// TODO: make this singleton with thread safety
func GetPlayNextCommand(bot DvonnBot) *PlayNextCommand {
	return &PlayNextCommand{bot}
}

func (playNextCmd *PlayNextCommand) Execute() {
	playNextCmd.bot.playNextMove()
}
