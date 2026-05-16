package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Go(gs *game.GameState, cmd parser.Command) {
	direction := cmd.DirectObject
	room := world.Rooms[gs.CurrentRoom]

	nextRoom, ok := room.Exits[direction]
	if !ok {
		output.Println("You can't go that way")
		return
	}

	gs.CurrentRoom = nextRoom
	Look(gs, parser.Command{})
}
