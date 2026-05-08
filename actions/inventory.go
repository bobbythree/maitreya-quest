package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
)

func Inventory(gs *game.GameState, cmd parser.Command) {
	if len(gs.Player.Inventory) == 0 {
		fmt.Println("You are carrying nothing.")
		return
	}

	fmt.Println("You are carrying:")

	for _, item := range gs.Player.Inventory {
		fmt.Println("-", item)
	}
}
