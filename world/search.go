package world

import "github.com/bobbythree/maitreya-quest/game"

// helper for searching nested containers

func findObjectRecursive(obj Object, target string) (Object, bool) {
	for _, childID := range obj.Contains {

		child := Objects[childID]

		if child.ID == target {
			return child, true
		}

		if child.Container {
			if !child.Openable || child.Open {

				found, ok := findObjectRecursive(child, target)

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

		obj := Objects[objID]

		if obj.ID == target {
			return obj, true
		}

		if obj.Container {
			if !obj.Openable || obj.Open {

				found, ok := findObjectRecursive(obj, target)

				if ok {
					return found, true
				}
			}
		}
	}

	// search inventory

	for _, objID := range gs.Player.Inventory {

		obj := Objects[objID]

		if obj.ID == target {
			return obj, true
		}
	}

	return Object{}, false
}
