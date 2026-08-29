package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Go(gs *game.GameState, cmd parser.Command) string {
	direction := cmd.DirectObject
	room := world.Rooms[gs.CurrentRoom]

	nextRoom, ok := room.Exits[direction]
	if !ok {
		return "You can't go that way"
	}

	// block access to the building until the security door is unlocked
	// TODO: get rid of this harcoded value later.
	if gs.CurrentRoom == "work_main" && direction == "north" && !gs.Flags["security_door_unlocked"] {
		return "The security door is locked."
	}

	gs.CurrentRoom = nextRoom
	next := world.Rooms[nextRoom]

	if !gs.VisitedRooms[nextRoom] {
		gs.VisitedRooms[nextRoom] = true

		if next.FirstVisitNarration != "" {
			return next.FirstVisitNarration + "\n\n" + next.Description
		}

		return next.Description
	}

	return ""
}
