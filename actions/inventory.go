package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
)

func Inventory(gs *game.GameState, cmd parser.Command) {
	if len(gs.Player.Inventory) == 0 {
		output.Println("You are carrying nothing.")
		return
	}

	output.Println("You are carrying:")

	for _, item := range gs.Player.Inventory {
		output.Println("-" + item)
	}
}
