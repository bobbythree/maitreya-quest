package workcomputer

import "github.com/bobbythree/maitreya-quest/game"

// Start begins interaction with the work computer.
func Start(gs *game.GameState) string {
	if !gs.BeginWorkComputer(game.WorkComputerState{
		Screen: "menu",
	}) {
		return "You're already busy."
	}

	return ""
}
