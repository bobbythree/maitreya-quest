package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Close(gs *game.GameState, cmd parser.Command) {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		fmt.Println("You don't see that.")
		return
	}

	if !obj.Openable {
		fmt.Println("You can't close that.")
		return
	}

	if !obj.Open {
		fmt.Println("It's already closed.")
		return
	}

	obj.Open = false

	world.Objects[obj.ID] = obj

	fmt.Printf("You close the %s.\n", obj.Name)
}
