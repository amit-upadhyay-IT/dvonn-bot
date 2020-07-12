package bot

import "github.com/amit-upadhyay-it/dvonn/dvonn"

type BotCommand interface {
	Execute(game *dvonn.DvonnGame) (bool, []string)
}


