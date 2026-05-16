package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Open(gs *game.GameState, cmd parser.Command) {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		output.Println("You don't see that.")
		return
	}

	if !obj.Openable {
		output.Println("You can't open that.")
		return
	}

	if obj.Open {
		output.Println("It's already open.")
		return
	}

	if obj.Locked {
		output.Println("It's locked.")
		return
	}

	obj.Open = true

	world.Objects[obj.ID] = obj

	output.Printf("You open the %s.\n", obj.Name)
}
