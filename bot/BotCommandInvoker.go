package bot

type BotCommandInvoker struct {
	command BotCommand
}

func GetBotCommandInvoker(command BotCommand) *BotCommandInvoker {
	return &BotCommandInvoker{command}
}

func (invoker *BotCommandInvoker) PlayAMove() {
	invoker.command.Execute()
}
