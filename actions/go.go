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

	gs.CurrentRoom = nextRoom

	// CurrentRoom has changed, so Look now gets the new room's description.
	return Look(gs, parser.Command{})
}
