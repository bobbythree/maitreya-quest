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

	state, ok := gs.ObjectStates[obj.ID]
	if !ok {
		return "You don't see that."
	}

	if state.Parent == "inventory" {
		return "You already have it."
	}

	if !obj.Portable {
		return "You can't take that."
	}

	gs.Player.Inventory = append(gs.Player.Inventory, obj.ID)

	if state.Parent != "" {

		parentState, ok := gs.ObjectStates[state.Parent]
		if ok {
			for i, childID := range parentState.Contains {
				if childID == obj.ID {

					parentState.Contains = append(
						parentState.Contains[:i],
						parentState.Contains[i+1:]...,
					)

					break
				}
			}
		}
	}

	state.Parent = "inventory"

	return "Taken."
}
