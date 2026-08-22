package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Get(gs *game.GameState, cmd parser.Command) string {
	directObject := cmd.DirectObject

	obj, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return "You don't see that."
	}

	if !obj.Portable {
		return "You can't take that."
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

	return "Taken."
}
