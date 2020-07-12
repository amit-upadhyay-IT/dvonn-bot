package bot

import "github.com/amit-upadhyay-it/dvonn/dvonn"

type DvonnBot interface {
	playNextMove(game *dvonn.DvonnGame) (bool, []string)
}
