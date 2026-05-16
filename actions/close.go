package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Close(gs *game.GameState, cmd parser.Command) {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		output.Println("You don't see that.")
		return
	}

	if !obj.Openable {
		output.Println("You can't close that.")
		return
	}

	if !obj.Open {
		output.Println("It's already closed.")
		return
	}

	obj.Open = false

	world.Objects[obj.ID] = obj

	output.Printf("You close the %s.\n", obj.Name)
}
