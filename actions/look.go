package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Look(gs *game.GameState, cmd parser.Command) {
	room := world.Rooms[gs.CurrentRoom]
	directObject := cmd.DirectObject

	if directObject == "" {
		output.Println(room.Description)
		return
	}

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		output.Println("You don't see that")
		return
	}

	if obj.Openable {
		if obj.Open {
			output.Println(obj.OpenDescription)
		} else {
			output.Println(obj.ClosedDescription)
		}
	} else {
		output.Println(obj.Description)
	}

	if obj.Container && obj.Open {
		if len(obj.Contains) > 0 {
			output.Println("You see:")
			for _, insideID := range obj.Contains {
				inside := world.Objects[insideID]
				output.Println("- " + inside.Name)
			}
		} else {
			output.Println("It's empty")
		}
	}
}
