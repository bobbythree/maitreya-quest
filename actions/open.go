package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Open(gs *game.GameState, cmd parser.Command) string {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return "You don't see that."
	}

	if !obj.Openable {
		return "You can't open that."
	}

	if obj.Open {
		return "It's already open."
	}

	if obj.Locked {
		return "It's locked"
	}

	obj.Open = true

	world.Objects[obj.ID] = obj

	return fmt.Sprintf("You open the %s.", obj.Name)
}
