package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Open(gs *game.GameState, cmd parser.Command) {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		fmt.Println("You don't see that.")
		return
	}

	if !obj.Openable {
		fmt.Println("You can't open that.")
		return
	}

	if obj.Open {
		fmt.Println("It's already open.")
		return
	}

	if obj.Locked {
		fmt.Println("It's locked.")
		return
	}

	obj.Open = true

	world.Objects[obj.ID] = obj

	fmt.Printf("You open the %s.\n", obj.Name)
}
