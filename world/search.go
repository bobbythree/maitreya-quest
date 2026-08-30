package world

import "github.com/bobbythree/maitreya-quest/game"

// helper for searching nested containers

func findObjectRecursive(gs *game.GameState, obj Object, target string) (Object, bool) {
	state, ok := gs.ObjectStates[obj.ID]
	if !ok {
		return Object{}, false
	}

	for _, childID := range state.Contains {

		child, ok := Objects[childID]
		if !ok {
			continue
		}

		if child.Name == target {
			return child, true
		}

		if child.Container {
			childState, ok := gs.ObjectStates[child.ID]
			if !ok {
				continue
			}

			if !child.Openable || childState.Open {

				found, ok := findObjectRecursive(gs, child, target)

				if ok {
					return found, true
				}
			}
		}
	}

	return Object{}, false
}

// searches visible room objects and player inventory

func FindVisibleObject(gs *game.GameState, target string) (Object, bool) {
	room := Rooms[gs.CurrentRoom]

	// search room objects

	for _, objID := range room.Objects {

		obj, ok := Objects[objID]
		if !ok {
			continue
		}

		if obj.Name == target {
			return obj, true
		}

		if obj.Container {
			state, ok := gs.ObjectStates[obj.ID]
			if !ok {
				continue
			}

			if !obj.Openable || state.Open {

				found, ok := findObjectRecursive(gs, obj, target)

				if ok {
					return found, true
				}
			}
		}
	}

	// search inventory

	for _, objID := range gs.Player.Inventory {

		obj, ok := Objects[objID]
		if !ok {
			continue
		}

		if obj.Name == target {
			return obj, true
		}
	}

	return Object{}, false
}
