package workcomputer

import "github.com/bobbythree/maitreya-quest/game"

func Start(gs *game.GameState) string {
	gs.WorkComputer = &game.WorkComputerState{
		Cursor: 0,
	}

	return ""
}
