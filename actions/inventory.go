package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
)

func Inventory(gs *game.GameState) {
	if len(gs.Inventory) == 0 {
		fmt.Println("You are carrying nothing.")
		return
	}

	fmt.Println("You are carrying:")

	for _, item := range gs.Inventory {
		fmt.Println("-", item)
	}
}
