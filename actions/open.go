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

	state, ok := gs.ObjectStates[obj.ID]
	if !ok {
		return "You can't open that."
	}

	if state.Open {
		return "It's already open."
	}

	if state.Locked {
		return "It's locked"
	}

	state.Open = true

	return fmt.Sprintf("You open the %s.", obj.Name)
}
