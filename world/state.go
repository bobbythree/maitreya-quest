package world

import "github.com/bobbythree/maitreya-quest/game"

// InitializeObjectStates creates fresh runtime state for every object.
func InitializeObjectStates(gs *game.GameState) {
	for id, obj := range Objects {
		gs.ObjectStates[id] = &game.ObjectState{
			Open:     obj.Open,
			Locked:   obj.Locked,
			Parent:   obj.Parent,
			Contains: append([]string(nil), obj.Contains...),
		}
	}
}
