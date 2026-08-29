package workcomputer

import "github.com/bobbythree/maitreya-quest/game"

// Start begins interaction with the work computer.
func Start(gs *game.GameState) string {
	gs.WorkComputer = &game.WorkComputerState{
		Screen: "menu",
	}

	return ""
}
