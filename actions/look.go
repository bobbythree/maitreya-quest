package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Look(gs *game.GameState, cmd parser.Command) {
	room := world.Rooms[gs.CurrentRoom]
	directObject := cmd.DirectObject

	if directObject == "" {
		fmt.Println(room.Description)
		return
	}

	obj, ok := world.FindVisibleObject(room, directObject)

	if !ok {
		fmt.Println("You don't see that")
		return
	}

	if obj.Openable {
		if obj.Open {
			fmt.Println(obj.OpenDescription)
		} else {
			fmt.Println(obj.ClosedDescription)
		}
	} else {
		fmt.Println(obj.Description)
	}

	if obj.Container && obj.Open {
		if len(obj.Contains) > 0 {
			fmt.Println("You see:")
			for _, insideID := range obj.Contains {
				inside := world.Objects[insideID]
				fmt.Println("- ", inside.Name)
			}
		} else {
			fmt.Println("It's empty")
		}
	}
}
