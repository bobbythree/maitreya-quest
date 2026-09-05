package oldcomputer

import "github.com/bobbythree/maitreya-quest/game"

const destroyedFlag = "old_computer_destroyed"

// Use describes the result of trying to use the old computer on its own.
func Use(gs *game.GameState) string {
	if gs.Flags[destroyedFlag] {
		return "The old computer is dead. It won't turn on again."
	}

	return "The computer looks like it's 100 years old. It won't turn on."
}

// Start powers on the old computer using the inserted thumb drive.
func Start(gs *game.GameState) string {
	if gs.Flags[destroyedFlag] {
		return "The old computer is dead. It won't turn on again."
	}

	if !gs.BeginOldComputer(game.OldComputerState{Screen: "loading"}) {
		return "You're already busy."
	}

	return ""
}
