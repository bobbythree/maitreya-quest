package actions

import (
	"strings"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Look(gs *game.GameState, cmd parser.Command) string {
	room := world.Rooms[gs.CurrentRoom]
	directObject := cmd.DirectObject
	result := strings.Builder{}

	if directObject == "" {
		return room.Description
	}

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return "You don't see that."
	}

	state, ok := gs.ObjectStates[obj.ID]
	if !ok {
		return "You don't see that."
	}

	if obj.Openable {
		if state.Open {
			result.WriteString(obj.OpenDescription)
		} else {
			result.WriteString(obj.ClosedDescription)
		}
	} else {
		result.WriteString(obj.Description)
	}

	if obj.Container && state.Open {
		if len(state.Contains) > 0 {
			result.WriteString("\nYou see:")
			for _, insideID := range state.Contains {
				inside := world.Objects[insideID]
				result.WriteString("\n- " + inside.Name)
			}
		} else {
			result.WriteString("\nIt's empty")
		}
	}

	return result.String()
}
