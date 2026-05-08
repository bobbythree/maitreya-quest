package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Get(gs *game.GameState, cmd parser.Command) {
	noun := cmd.Noun

	room := world.Rooms[gs.CurrentRoom]

	obj, ok := world.FindVisibleObject(room, noun)

	if !ok {
		fmt.Println("You don't see that.")
		return
	}

	if !obj.Portable {
		fmt.Println("You can't take that.")
		return
	}

	gs.Player.Inventory = append(gs.Player.Inventory, obj.ID)

	if obj.Parent != "" {

		parent := world.Objects[obj.Parent]

		for i, childID := range parent.Contains {
			if childID == obj.ID {

				parent.Contains = append(
					parent.Contains[:i],
					parent.Contains[i+1:]...,
				)

				world.Objects[parent.ID] = parent

				break
			}
		}
	}

	obj.Parent = "inventory"
	world.Objects[obj.ID] = obj

	fmt.Println("Taken.")
}
