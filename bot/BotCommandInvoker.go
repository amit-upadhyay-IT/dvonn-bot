package bot

import "github.com/amit-upadhyay-it/dvonn/dvonn"

type BotCommandInvoker struct {
	command BotCommand
}

func GetBotCommandInvoker(command BotCommand) BotCommandInvoker {
	return BotCommandInvoker{command}
}

func (invoker BotCommandInvoker) PlayAMove(game *dvonn.DvonnGame) (bool, []string) {
	return invoker.command.Execute(game)
}
