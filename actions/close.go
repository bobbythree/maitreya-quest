package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Close(gs *game.GameState, cmd parser.Command) string {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return "You don't see that."
	}

	if !obj.Openable {
		return "You can't close that."
	}

	state, ok := gs.ObjectStates[obj.ID]
	if !ok {
		return "You can't close that."
	}

	if !state.Open {
		return "It's already closed."
	}

	state.Open = false

	return fmt.Sprintf("You close the %s.", obj.Name)
}
